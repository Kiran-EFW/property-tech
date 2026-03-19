package service

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// ProjectFilters holds the query filters for listing projects.
type ProjectFilters struct {
	Status   string  `query:"status"`
	City     string  `query:"city"`
	MinPrice float64 `query:"min_price"`
	MaxPrice float64 `query:"max_price"`
	Page     int     `query:"page"`
	Limit    int     `query:"limit"`
}

// CreateProjectInput holds the fields needed to create a project.
type CreateProjectInput struct {
	Name          string   `json:"name" validate:"required"`
	RERANumber    string   `json:"rera_number" validate:"required"`
	BuilderID     string   `json:"builder_id" validate:"required,uuid"`
	Description   *string  `json:"description"`
	CarpetAreaMin *float64 `json:"carpet_area_min"`
	CarpetAreaMax *float64 `json:"carpet_area_max"`
	PriceMin      *float64 `json:"price_min"`
	PriceMax      *float64 `json:"price_max"`
	Location      string   `json:"location"`
	Address       *string  `json:"address"`
	City          *string  `json:"city"`
	State         *string  `json:"state"`
	Pincode       *string  `json:"pincode"`
}

// UpdateProjectInput holds the fields that can be updated on a project.
type UpdateProjectInput struct {
	Name          *string          `json:"name"`
	Description   *string          `json:"description"`
	CarpetAreaMin *float64         `json:"carpet_area_min"`
	CarpetAreaMax *float64         `json:"carpet_area_max"`
	PriceMin      *float64         `json:"price_min"`
	PriceMax      *float64         `json:"price_max"`
	Location      *string          `json:"location"`
	Address       *string          `json:"address"`
	City          *string          `json:"city"`
	State         *string          `json:"state"`
	Pincode       *string          `json:"pincode"`
	Status        *string          `json:"status"`
}

// ProjectRepository defines the database operations required by ProjectService.
type ProjectRepository interface {
	ListProjects(ctx context.Context, filters ProjectFilters) ([]*domain.Project, int, error)
	GetProjectBySlug(ctx context.Context, slug string) (*domain.Project, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	CreateProject(ctx context.Context, project *domain.Project) error
	UpdateProject(ctx context.Context, project *domain.Project) error
	GetProjectUnits(ctx context.Context, projectID uuid.UUID) ([]*domain.ProjectUnit, error)
	GetDueDiligence(ctx context.Context, projectID uuid.UUID) (*domain.DueDiligenceReport, error)
}

// ProjectService handles project business logic.
type ProjectService struct {
	repo ProjectRepository
}

// NewProjectService creates a new ProjectService.
func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// List returns a filtered, paginated list of projects.
func (s *ProjectService) List(ctx context.Context, filters ProjectFilters) ([]*domain.Project, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	projects, total, err := s.repo.ListProjects(ctx, filters)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list projects: %w", err)
	}

	return projects, total, nil
}

// GetBySlug returns a project by its URL slug.
func (s *ProjectService) GetBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	project, err := s.repo.GetProjectBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("project not found: %w", err)
	}
	return project, nil
}

// reraRegex validates Indian RERA registration numbers.
// Format varies by state but generally: STATE/RERA/ALPHA-NUMERIC
var reraRegex = regexp.MustCompile(`^[A-Z]{2,5}/\w+/\w+`)

// Create validates the input, generates a slug, and persists a new project.
func (s *ProjectService) Create(ctx context.Context, input CreateProjectInput) (*domain.Project, error) {
	// Validate RERA number format.
	if !reraRegex.MatchString(input.RERANumber) {
		return nil, fmt.Errorf("invalid RERA number format: %s", input.RERANumber)
	}

	builderID, err := uuid.Parse(input.BuilderID)
	if err != nil {
		return nil, fmt.Errorf("invalid builder_id: %w", err)
	}

	slug := generateSlug(input.Name)

	project := &domain.Project{
		ID:            uuid.New(),
		Name:          input.Name,
		Slug:          slug,
		RERANumber:    input.RERANumber,
		BuilderID:     builderID,
		Description:   input.Description,
		CarpetAreaMin: input.CarpetAreaMin,
		CarpetAreaMax: input.CarpetAreaMax,
		PriceMin:      input.PriceMin,
		PriceMax:      input.PriceMax,
		Location:      input.Location,
		Address:       input.Address,
		City:          input.City,
		State:         input.State,
		Pincode:       input.Pincode,
		Status:        domain.ProjectStatusDraft,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.CreateProject(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	log.Info().Str("project_id", project.ID.String()).Str("slug", slug).Msg("project created")
	return project, nil
}

// Update applies partial updates to an existing project.
func (s *ProjectService) Update(ctx context.Context, id uuid.UUID, input UpdateProjectInput) (*domain.Project, error) {
	project, err := s.repo.GetProjectByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("project not found: %w", err)
	}

	if input.Name != nil {
		project.Name = *input.Name
		project.Slug = generateSlug(*input.Name)
	}
	if input.Description != nil {
		project.Description = input.Description
	}
	if input.CarpetAreaMin != nil {
		project.CarpetAreaMin = input.CarpetAreaMin
	}
	if input.CarpetAreaMax != nil {
		project.CarpetAreaMax = input.CarpetAreaMax
	}
	if input.PriceMin != nil {
		project.PriceMin = input.PriceMin
	}
	if input.PriceMax != nil {
		project.PriceMax = input.PriceMax
	}
	if input.Location != nil {
		project.Location = *input.Location
	}
	if input.Address != nil {
		project.Address = input.Address
	}
	if input.City != nil {
		project.City = input.City
	}
	if input.State != nil {
		project.State = input.State
	}
	if input.Pincode != nil {
		project.Pincode = input.Pincode
	}
	if input.Status != nil {
		project.Status = domain.ProjectStatus(*input.Status)
	}
	project.UpdatedAt = time.Now()

	if err := s.repo.UpdateProject(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	log.Info().Str("project_id", id.String()).Msg("project updated")
	return project, nil
}

// GetInventory returns all units for a project.
func (s *ProjectService) GetInventory(ctx context.Context, projectID uuid.UUID) ([]*domain.ProjectUnit, error) {
	units, err := s.repo.GetProjectUnits(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}
	return units, nil
}

// GetDueDiligence returns the due-diligence report for a project.
func (s *ProjectService) GetDueDiligence(ctx context.Context, projectID uuid.UUID) (*domain.DueDiligenceReport, error) {
	report, err := s.repo.GetDueDiligence(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get due-diligence report: %w", err)
	}
	return report, nil
}

// generateSlug creates a URL-friendly slug from a project name.
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove non-alphanumeric characters (except hyphens).
	re := regexp.MustCompile(`[^a-z0-9\-]`)
	slug = re.ReplaceAllString(slug, "")
	// Collapse multiple hyphens.
	re = regexp.MustCompile(`-+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}
