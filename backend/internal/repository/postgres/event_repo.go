package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
)

// EventRepo implements service.EventRepository using raw SQL against a pgx pool.
type EventRepo struct {
	pool *pgxpool.Pool
}

// NewEventRepo creates a new EventRepo backed by the given connection pool.
func NewEventRepo(pool *pgxpool.Pool) *EventRepo {
	return &EventRepo{pool: pool}
}

// CreateEvent inserts an immutable audit-log event into the events table.
func (r *EventRepo) CreateEvent(ctx context.Context, event *domain.Event) error {
	query := `
		INSERT INTO events (id, actor_id, actor_role, action, entity_type, entity_id, payload, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.pool.Exec(ctx, query,
		event.ID,
		event.ActorID,
		event.ActorRole,
		event.Action,
		event.EntityType,
		event.EntityID,
		event.Payload,
		event.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	return nil
}
