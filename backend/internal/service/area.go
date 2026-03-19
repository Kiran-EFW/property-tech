package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// AreaFilters holds the query filters for listing areas.
type AreaFilters struct {
	City  string `query:"city"`
	State string `query:"state"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

// CreateAreaInput holds the fields needed to create an area.
type CreateAreaInput struct {
	Name        string  `json:"name" validate:"required"`
	City        string  `json:"city" validate:"required"`
	State       string  `json:"state" validate:"required"`
	Description *string `json:"description"`
	Location    string  `json:"location"`
}

// AreaRepository defines the database operations required by AreaService.
type AreaRepository interface {
	CreateArea(ctx context.Context, area *domain.Area) error
	GetAreaBySlug(ctx context.Context, slug string) (*domain.Area, error)
	ListAreas(ctx context.Context, filters AreaFilters) ([]*domain.Area, int, error)
}

// AreaService handles area/micro-market business logic.
type AreaService struct {
	repo AreaRepository
}

// NewAreaService creates a new AreaService.
func NewAreaService(repo AreaRepository) *AreaService {
	return &AreaService{repo: repo}
}

// List returns a filtered, paginated list of areas.
func (s *AreaService) List(ctx context.Context, filters AreaFilters) ([]*domain.Area, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	areas, total, err := s.repo.ListAreas(ctx, filters)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list areas: %w", err)
	}

	return areas, total, nil
}

// GetBySlug returns an area by its URL slug.
func (s *AreaService) GetBySlug(ctx context.Context, slug string) (*domain.Area, error) {
	area, err := s.repo.GetAreaBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("area not found: %w", err)
	}
	return area, nil
}

// Create validates the input and persists a new area.
func (s *AreaService) Create(ctx context.Context, input CreateAreaInput) (*domain.Area, error) {
	slug := generateSlug(input.Name + " " + input.City)

	area := &domain.Area{
		ID:          uuid.New(),
		Name:        input.Name,
		Slug:        slug,
		City:        input.City,
		State:       input.State,
		Description: input.Description,
		Location:    input.Location,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateArea(ctx, area); err != nil {
		return nil, fmt.Errorf("failed to create area: %w", err)
	}

	log.Info().Str("area_id", area.ID.String()).Str("slug", slug).Msg("area created")
	return area, nil
}
