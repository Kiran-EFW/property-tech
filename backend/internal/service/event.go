package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// EventRepository defines the database operations required by EventService.
type EventRepository interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
}

// EventService records audit trail entries for significant actions.
type EventService struct {
	repo EventRepository
}

// NewEventService creates a new EventService.
func NewEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

// Log creates an immutable audit log entry.
func (s *EventService) Log(ctx context.Context, actorID uuid.UUID, actorRole, action, entityType string, entityID uuid.UUID, payload map[string]interface{}) error {
	var payloadJSON json.RawMessage
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal event payload: %w", err)
		}
		payloadJSON = data
	}

	event := &domain.Event{
		ID:         uuid.New(),
		ActorID:    actorID,
		ActorRole:  actorRole,
		Action:     action,
		EntityType: entityType,
		EntityID:   entityID,
		Payload:    payloadJSON,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.CreateEvent(ctx, event); err != nil {
		log.Error().Err(err).
			Str("action", action).
			Str("entity_type", entityType).
			Str("entity_id", entityID.String()).
			Msg("failed to log event")
		return fmt.Errorf("failed to log event: %w", err)
	}

	log.Debug().
		Str("action", action).
		Str("entity_type", entityType).
		Str("entity_id", entityID.String()).
		Str("actor_id", actorID.String()).
		Msg("event logged")

	return nil
}
