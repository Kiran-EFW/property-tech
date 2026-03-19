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

// CreateVisitInput holds the fields needed to create a site visit.
type CreateVisitInput struct {
	LeadID    string `json:"lead_id" validate:"required,uuid"`
	ProjectID string `json:"project_id" validate:"required,uuid"`
	VisitedAt string `json:"visited_at"` // RFC3339 format; defaults to now
}

// SubmitFeedbackInput holds the feedback payload for a site visit.
type SubmitFeedbackInput struct {
	Feedback string   `json:"feedback"`
	Photos   []string `json:"photos"`
	Duration *int     `json:"duration"` // minutes
	Rating   *int     `json:"rating"`   // 1-5
}

// VisitFilters holds the query filters for listing visits.
type VisitFilters struct {
	ProjectID string `query:"project_id"`
	LeadID    string `query:"lead_id"`
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
}

// VisitRepository defines the database operations required by VisitService.
type VisitRepository interface {
	CreateVisit(ctx context.Context, visit *domain.SiteVisit) error
	GetVisitByID(ctx context.Context, id uuid.UUID) (*domain.SiteVisit, error)
	ListVisits(ctx context.Context, filters VisitFilters) ([]*domain.SiteVisit, int, error)
	ListVisitsByAgent(ctx context.Context, agentID uuid.UUID, filters VisitFilters) ([]*domain.SiteVisit, int, error)
	UpdateVisit(ctx context.Context, visit *domain.SiteVisit) error
}

// VisitService handles site visit business logic.
type VisitService struct {
	repo VisitRepository
}

// NewVisitService creates a new VisitService.
func NewVisitService(repo VisitRepository) *VisitService {
	return &VisitService{repo: repo}
}

// Create creates a new site visit record.
func (s *VisitService) Create(ctx context.Context, agentID uuid.UUID, input CreateVisitInput) (*domain.SiteVisit, error) {
	leadID, err := uuid.Parse(input.LeadID)
	if err != nil {
		return nil, fmt.Errorf("invalid lead_id: %w", err)
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("invalid project_id: %w", err)
	}

	visitedAt := time.Now()
	if input.VisitedAt != "" {
		parsed, err := time.Parse(time.RFC3339, input.VisitedAt)
		if err == nil {
			visitedAt = parsed
		}
	}

	visit := &domain.SiteVisit{
		ID:        uuid.New(),
		LeadID:    leadID,
		AgentID:   agentID,
		ProjectID: projectID,
		VisitedAt: visitedAt,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateVisit(ctx, visit); err != nil {
		return nil, fmt.Errorf("failed to create visit: %w", err)
	}

	log.Info().Str("visit_id", visit.ID.String()).Str("agent_id", agentID.String()).Msg("site visit created")
	return visit, nil
}

// SubmitFeedback adds feedback, photos, and rating to a site visit.
func (s *VisitService) SubmitFeedback(ctx context.Context, visitID uuid.UUID, input SubmitFeedbackInput) error {
	visit, err := s.repo.GetVisitByID(ctx, visitID)
	if err != nil {
		return fmt.Errorf("visit not found: %w", err)
	}

	if input.Feedback != "" {
		visit.Feedback = &input.Feedback
	}
	if input.Duration != nil {
		visit.Duration = input.Duration
	}
	if input.Rating != nil {
		visit.Rating = input.Rating
	}
	if len(input.Photos) > 0 {
		photosJSON, err := json.Marshal(input.Photos)
		if err != nil {
			return fmt.Errorf("failed to marshal photos: %w", err)
		}
		visit.Photos = photosJSON
	}

	if err := s.repo.UpdateVisit(ctx, visit); err != nil {
		return fmt.Errorf("failed to update visit: %w", err)
	}

	log.Info().Str("visit_id", visitID.String()).Msg("visit feedback submitted")
	return nil
}

// List returns visits, scoped by role. Agents see only their own visits.
func (s *VisitService) List(ctx context.Context, filters VisitFilters, role string, userID uuid.UUID) ([]*domain.SiteVisit, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	if role == "admin" || role == "super_admin" {
		return s.repo.ListVisits(ctx, filters)
	}

	return s.repo.ListVisitsByAgent(ctx, userID, filters)
}
