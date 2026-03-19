package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

// Task type constants for all background jobs.
const (
	TaskLeadEscalation  = "task:lead_escalation"
	TaskSendNotification = "task:send_notification"
	TaskRERARecheck      = "task:rera_recheck"
)

// Worker manages background task processing using Asynq.
type Worker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	client *asynq.Client

	// Dependencies injected from main.
	leadEscalation  *LeadEscalationHandler
	notifications   *NotificationHandler
}

// WorkerConfig holds configuration for the worker.
type WorkerConfig struct {
	RedisAddr   string
	Concurrency int
}

// NewWorker creates a new Worker with the given configuration and dependencies.
func NewWorker(cfg WorkerConfig, leadEscalation *LeadEscalationHandler, notifications *NotificationHandler) *Worker {
	redisOpt := asynq.RedisClientOpt{Addr: cfg.RedisAddr}

	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: cfg.Concurrency,
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().
				Err(err).
				Str("task_type", task.Type()).
				Msg("worker: task failed")
		}),
	})

	client := asynq.NewClient(redisOpt)

	return &Worker{
		server:          server,
		mux:             asynq.NewServeMux(),
		client:          client,
		leadEscalation:  leadEscalation,
		notifications:   notifications,
	}
}

// RegisterTasks registers all async task handlers on the mux.
func (w *Worker) RegisterTasks() {
	w.mux.HandleFunc(TaskLeadEscalation, w.leadEscalation.HandleLeadEscalation)
	w.mux.HandleFunc(TaskSendNotification, w.notifications.HandleSendNotification)
	w.mux.HandleFunc(TaskRERARecheck, w.handleRERARecheck)

	log.Info().Msg("worker: all task handlers registered")
}

// Start begins processing tasks. This blocks until the server is shut down.
func (w *Worker) Start() error {
	w.RegisterTasks()

	log.Info().Msg("worker: starting task processing")
	if err := w.server.Start(w.mux); err != nil {
		return fmt.Errorf("worker: failed to start: %w", err)
	}
	return nil
}

// Shutdown gracefully shuts down the worker.
func (w *Worker) Shutdown() {
	w.server.Shutdown()
	if err := w.client.Close(); err != nil {
		log.Error().Err(err).Msg("worker: failed to close client")
	}
	log.Info().Msg("worker: shut down")
}

// Client returns the Asynq client for enqueuing tasks.
func (w *Worker) Client() *asynq.Client {
	return w.client
}

// handleRERARecheck processes RERA re-verification tasks.
func (w *Worker) handleRERARecheck(ctx context.Context, task *asynq.Task) error {
	log.Info().
		Str("task_type", task.Type()).
		Msg("worker: processing RERA recheck")

	// TODO: Implement RERA re-verification logic.
	// 1. Fetch project by ID from task payload
	// 2. Query the RERA portal API for current status
	// 3. Update the due-diligence report with fresh data
	// 4. Log the event

	return nil
}
