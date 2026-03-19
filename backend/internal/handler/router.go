package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/proptech/backend/internal/config"
	"github.com/proptech/backend/internal/middleware"
)

// Handlers aggregates all HTTP handler instances.
type Handlers struct {
	Auth       *AuthHandler
	Project    *ProjectHandler
	Lead       *LeadHandler
	Agent      *AgentHandler
	Visit      *VisitHandler
	Commission *CommissionHandler
	Area       *AreaHandler
	Builder    *BuilderHandler
	Search     *SearchHandler
	Media      *MediaHandler
}

// Middleware aggregates the middleware dependencies needed by the router.
type Middleware struct {
	JWTAuth fiber.Handler
}

// NewMiddleware creates a new Middleware bundle from the config.
func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{
		JWTAuth: middleware.NewJWTAuth(cfg),
	}
}

// SetupRoutes registers all API routes on the Fiber app.
func SetupRoutes(app *fiber.App, h *Handlers, mw *Middleware) {
	api := app.Group("/api/v1")

	// -----------------------------------------------------------------------
	// Auth routes (public)
	// -----------------------------------------------------------------------
	auth := api.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.RequestOTP)
	auth.Post("/verify", h.Auth.VerifyOTP)
	auth.Post("/refresh", h.Auth.RefreshToken)

	// Auth routes (authenticated)
	authProtected := auth.Group("", mw.JWTAuth)
	authProtected.Get("/me", h.Auth.GetMe)
	authProtected.Put("/profile", h.Auth.UpdateProfile)

	// -----------------------------------------------------------------------
	// Project routes
	// -----------------------------------------------------------------------
	projects := api.Group("/projects")
	// Public endpoints.
	projects.Get("/", h.Project.List)
	projects.Get("/:slug", h.Project.GetBySlug)

	// Authenticated project endpoints.
	projectsAuth := projects.Group("", mw.JWTAuth)
	projectsAuth.Get("/:id/inventory", h.Project.GetInventory)
	projectsAuth.Get("/:id/due-diligence", h.Project.GetDueDiligence)

	// Admin-only project endpoints.
	projectsAdmin := projects.Group("", mw.JWTAuth, middleware.RequireAdmin())
	projectsAdmin.Post("/", h.Project.Create)
	projectsAdmin.Put("/:id", h.Project.Update)

	// -----------------------------------------------------------------------
	// Lead routes
	// -----------------------------------------------------------------------
	leads := api.Group("/leads")
	// Public: investor submits interest.
	leads.Post("/", h.Lead.Create)

	// Authenticated lead endpoints.
	leadsAuth := leads.Group("", mw.JWTAuth)
	leadsAuth.Get("/", h.Lead.List)
	leadsAuth.Get("/:id", h.Lead.GetByID)
	leadsAuth.Put("/:id/status", h.Lead.UpdateStatus)
	leadsAuth.Post("/:id/notes", h.Lead.AddNote)

	// Admin-only lead endpoints.
	leadsAdmin := leads.Group("", mw.JWTAuth, middleware.RequireAdmin())
	leadsAdmin.Put("/:id/assign", h.Lead.Assign)

	// -----------------------------------------------------------------------
	// Site Visit routes (authenticated)
	// -----------------------------------------------------------------------
	visits := api.Group("/visits", mw.JWTAuth)
	visits.Post("/", h.Visit.Create)
	visits.Put("/:id/feedback", h.Visit.SubmitFeedback)
	visits.Get("/", h.Visit.List)

	// -----------------------------------------------------------------------
	// Agent routes
	// -----------------------------------------------------------------------
	agents := api.Group("/agents", mw.JWTAuth)
	agents.Post("/register", h.Agent.Register)
	agents.Get("/:id/performance", h.Agent.GetPerformance)

	// Admin-only agent endpoints.
	agentsAdmin := api.Group("/agents", mw.JWTAuth, middleware.RequireAdmin())
	agentsAdmin.Get("/", h.Agent.List)
	agentsAdmin.Put("/:id/tier", h.Agent.UpdateTier)

	// -----------------------------------------------------------------------
	// Commission routes (authenticated)
	// -----------------------------------------------------------------------
	commissions := api.Group("/commissions", mw.JWTAuth)
	commissions.Get("/", h.Commission.List)
	commissions.Get("/summary", h.Commission.GetSummary)

	// Admin-only commission endpoints.
	commissionsAdmin := api.Group("/commissions", mw.JWTAuth, middleware.RequireAdmin())
	commissionsAdmin.Post("/:id/approve", h.Commission.Approve)

	// -----------------------------------------------------------------------
	// Builder routes (authenticated, builder role)
	// -----------------------------------------------------------------------
	builders := api.Group("/builders", mw.JWTAuth, middleware.RequireRole("builder", "admin", "super_admin"))
	builders.Get("/me/projects", h.Builder.GetMyProjects)
	builders.Put("/me/inventory", h.Builder.UpdateInventory)
	builders.Get("/me/leads", h.Builder.GetMyLeads)

	// -----------------------------------------------------------------------
	// Area routes
	// -----------------------------------------------------------------------
	areas := api.Group("/areas")
	// Public endpoints.
	areas.Get("/", h.Area.List)
	areas.Get("/:slug", h.Area.GetBySlug)

	// Admin-only area endpoints.
	areasAdmin := areas.Group("", mw.JWTAuth, middleware.RequireAdmin())
	areasAdmin.Post("/", h.Area.Create)

	// -----------------------------------------------------------------------
	// Search routes (public)
	// -----------------------------------------------------------------------
	search := api.Group("/search")
	search.Get("/projects", h.Search.SearchProjects)
	search.Get("/areas", h.Search.SearchAreas)

	// -----------------------------------------------------------------------
	// Media routes (authenticated)
	// -----------------------------------------------------------------------
	media := api.Group("/media", mw.JWTAuth)
	media.Post("/upload-url", h.Media.GetUploadURL)
	media.Delete("/:id", h.Media.Delete)
}
