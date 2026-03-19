package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/adapter/sms"
	"github.com/proptech/backend/internal/adapter/storage"
	"github.com/proptech/backend/internal/config"
	"github.com/proptech/backend/internal/database"
	"github.com/proptech/backend/internal/handler"
	"github.com/proptech/backend/internal/middleware"
	pgrepo "github.com/proptech/backend/internal/repository/postgres"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/internal/worker"
	"github.com/proptech/backend/pkg/logger"
)

func main() {
	// Load configuration.
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Setup logger.
	logger.Setup(cfg.Environment)
	log.Info().Str("environment", cfg.Environment).Msg("starting proptech API")

	// Connect to PostgreSQL.
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to PostgreSQL")
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to ping PostgreSQL")
	}
	log.Info().Msg("connected to PostgreSQL")

	// Run database migrations if enabled.
	if cfg.AutoMigrate {
		if err := database.RunMigrations(cfg.DatabaseURL, cfg.MigrationsPath); err != nil {
			log.Error().Err(err).Msg("failed to run migrations")
		}
	}

	// Connect to Redis.
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse Redis URL")
	}
	redisClient := redis.NewClient(redisOpt)
	defer redisClient.Close()

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal().Err(err).Msg("failed to connect to Redis")
	}
	log.Info().Msg("connected to Redis")

	// -----------------------------------------------------------------------
	// Initialize adapters
	// -----------------------------------------------------------------------

	// SMS provider.
	var smsProvider sms.SMSProvider
	if cfg.TwilioAccountSID != "" {
		smsProvider = sms.NewTwilioProvider(cfg.TwilioAccountSID, cfg.TwilioAuthToken, cfg.TwilioFromNumber)
		log.Info().Msg("SMS provider: Twilio")
	} else {
		smsProvider = &sms.NoopSMSProvider{}
		log.Info().Msg("SMS provider: noop (dev mode)")
	}

	// Storage provider (Cloudflare R2).
	var storageProvider storage.StorageProvider
	if cfg.R2AccountID != "" && cfg.R2AccessKeyID != "" {
		r2, err := storage.NewR2Provider(cfg.R2AccountID, cfg.R2AccessKeyID, cfg.R2AccessKeySecret, cfg.R2Bucket, cfg.R2PublicURL)
		if err != nil {
			log.Error().Err(err).Msg("failed to initialize R2 storage provider")
		} else {
			storageProvider = r2
			log.Info().Msg("storage provider: Cloudflare R2")
		}
	}

	// Meilisearch client.
	var meiliClient meilisearch.ServiceManager
	if cfg.MeilisearchURL != "" {
		meiliClient = meilisearch.New(cfg.MeilisearchURL, meilisearch.WithAPIKey(cfg.MeilisearchKey))
		log.Info().Str("url", cfg.MeilisearchURL).Msg("Meilisearch client initialized")
	}

	// -----------------------------------------------------------------------
	// Initialize repositories
	// -----------------------------------------------------------------------

	authRepo := pgrepo.NewAuthRepo(pool)
	eventRepo := pgrepo.NewEventRepo(pool)
	projectRepo := pgrepo.NewProjectRepo(pool)
	leadRepo := pgrepo.NewLeadRepo(pool)
	agentRepo := pgrepo.NewAgentRepo(pool)
	visitRepo := pgrepo.NewVisitRepo(pool)
	commissionRepo := pgrepo.NewCommissionRepo(pool)
	areaRepo := pgrepo.NewAreaRepo(pool)

	log.Info().Msg("repositories initialized")

	// -----------------------------------------------------------------------
	// Initialize services
	// -----------------------------------------------------------------------

	// Event service (audit log) - used by other services.
	eventSvc := service.NewEventService(eventRepo)

	// Auth service.
	authSvc := service.NewAuthService(authRepo, redisClient, cfg)

	// Project service.
	projectSvc := service.NewProjectService(projectRepo)

	// Lead service.
	leadSvc := service.NewLeadService(leadRepo, eventSvc)

	// Agent service.
	agentSvc := service.NewAgentService(agentRepo, eventSvc)

	// Visit service.
	visitSvc := service.NewVisitService(visitRepo)

	// Commission service.
	commissionSvc := service.NewCommissionService(commissionRepo, eventSvc)

	// Area service.
	areaSvc := service.NewAreaService(areaRepo)

	_ = smsProvider

	// -----------------------------------------------------------------------
	// Initialize handlers
	// -----------------------------------------------------------------------

	handlers := &handler.Handlers{
		Auth:       handler.NewAuthHandler(authSvc),
		Project:    handler.NewProjectHandler(projectSvc),
		Lead:       handler.NewLeadHandler(leadSvc),
		Agent:      handler.NewAgentHandler(agentSvc),
		Visit:      handler.NewVisitHandler(visitSvc),
		Commission: handler.NewCommissionHandler(commissionSvc),
		Area:       handler.NewAreaHandler(areaSvc),
		Builder:    handler.NewBuilderHandler(projectSvc, leadSvc),
		Search:     handler.NewSearchHandler(meiliClient),
		Media:      handler.NewMediaHandler(storageProvider),
	}

	// -----------------------------------------------------------------------
	// Initialize workers
	// -----------------------------------------------------------------------

	// Extract Redis address from the URL for Asynq.
	redisAddr := extractRedisAddr(cfg.RedisURL)

	leadEscalationHandler := worker.NewLeadEscalationHandler(leadRepo)
	notificationHandler := worker.NewNotificationHandler(smsProvider)

	bgWorker := worker.NewWorker(worker.WorkerConfig{
		RedisAddr:   redisAddr,
		Concurrency: 10,
	}, leadEscalationHandler, notificationHandler)

	// Start background worker in a goroutine.
	go func() {
		if err := bgWorker.Start(); err != nil {
			log.Error().Err(err).Msg("background worker failed")
		}
	}()

	// -----------------------------------------------------------------------
	// Create Fiber app and register routes
	// -----------------------------------------------------------------------

	app := fiber.New(fiber.Config{
		AppName:      "PropTech API",
		ErrorHandler: defaultErrorHandler,
	})

	// Global middleware.
	app.Use(middleware.NewRequestID())
	app.Use(logger.FiberLogger())
	app.Use(middleware.NewCORS(cfg))
	app.Use(middleware.NewPrometheusMiddleware())

	// Health check.
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "proptech-api",
		})
	})

	// Metrics endpoint.
	app.Get("/metrics", middleware.MetricsHandler())

	// Register all API routes.
	mw := handler.NewMiddleware(cfg)
	handler.SetupRoutes(app, handlers, mw)

	// -----------------------------------------------------------------------
	// Graceful shutdown
	// -----------------------------------------------------------------------

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.ServerPort)
		log.Info().Str("addr", addr).Msg("listening")
		if err := app.Listen(addr); err != nil {
			log.Fatal().Err(err).Msg("server failed")
		}
	}()

	<-quit
	log.Info().Msg("shutting down server...")

	// Shutdown background worker.
	bgWorker.Shutdown()

	if err := app.Shutdown(); err != nil {
		log.Error().Err(err).Msg("error during shutdown")
	}
	log.Info().Msg("server stopped")
}

// defaultErrorHandler is the global Fiber error handler.
func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	log.Error().
		Err(err).
		Int("status", code).
		Str("path", c.Path()).
		Msg("unhandled error")

	return c.Status(code).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    "server_error",
			"message": message,
		},
	})
}

// extractRedisAddr parses a Redis URL and returns the host:port address for Asynq.
func extractRedisAddr(redisURL string) string {
	u, err := url.Parse(redisURL)
	if err != nil {
		return "localhost:6379"
	}
	host := u.Host
	if host == "" {
		return "localhost:6379"
	}
	return host
}
