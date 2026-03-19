package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/service"
)

// ProjectRepo implements service.ProjectRepository using raw SQL against a pgx pool.
type ProjectRepo struct {
	pool *pgxpool.Pool
}

// NewProjectRepo creates a new ProjectRepo backed by the given connection pool.
func NewProjectRepo(pool *pgxpool.Pool) *ProjectRepo {
	return &ProjectRepo{pool: pool}
}

// ListProjects returns a filtered, paginated list of projects along with the total count.
func (r *ProjectRepo) ListProjects(ctx context.Context, filters service.ProjectFilters) ([]*domain.Project, int, error) {
	query := `
		SELECT id, name, slug, rera_number, builder_id, description,
		       carpet_area_min, carpet_area_max, price_min, price_max,
		       COALESCE(ST_AsText(location), '') as location,
		       address, city, state, pincode, status,
		       amenities, media, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM projects`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, filters.Status)
		argIdx++
	}
	if filters.City != "" {
		conditions = append(conditions, fmt.Sprintf("city ILIKE $%d", argIdx))
		args = append(args, "%"+filters.City+"%")
		argIdx++
	}
	if filters.MinPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("price_min >= $%d", argIdx))
		args = append(args, filters.MinPrice)
		argIdx++
	}
	if filters.MaxPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("price_max <= $%d", argIdx))
		args = append(args, filters.MaxPrice)
		argIdx++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY created_at DESC"

	limit := filters.Limit
	if limit < 1 {
		limit = 20
	}
	page := filters.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list projects: %w", err)
	}
	defer rows.Close()

	var projects []*domain.Project
	var total int

	for rows.Next() {
		p := &domain.Project{}
		err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.RERANumber, &p.BuilderID, &p.Description,
			&p.CarpetAreaMin, &p.CarpetAreaMax, &p.PriceMin, &p.PriceMax,
			&p.Location,
			&p.Address, &p.City, &p.State, &p.Pincode, &p.Status,
			&p.Amenities, &p.Media, &p.CreatedAt, &p.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list projects scan: %w", err)
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list projects rows: %w", err)
	}

	return projects, total, nil
}

// GetProjectBySlug retrieves a single project by its URL-friendly slug.
func (r *ProjectRepo) GetProjectBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	query := `
		SELECT id, name, slug, rera_number, builder_id, description,
		       carpet_area_min, carpet_area_max, price_min, price_max,
		       COALESCE(ST_AsText(location), '') as location,
		       address, city, state, pincode, status,
		       amenities, media, created_at, updated_at
		FROM projects
		WHERE slug = $1`

	p := &domain.Project{}
	row := r.pool.QueryRow(ctx, query, slug)

	err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.RERANumber, &p.BuilderID, &p.Description,
		&p.CarpetAreaMin, &p.CarpetAreaMax, &p.PriceMin, &p.PriceMax,
		&p.Location,
		&p.Address, &p.City, &p.State, &p.Pincode, &p.Status,
		&p.Amenities, &p.Media, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get project by slug: %w", err)
	}

	return p, nil
}

// GetProjectByID retrieves a single project by its unique identifier.
func (r *ProjectRepo) GetProjectByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	query := `
		SELECT id, name, slug, rera_number, builder_id, description,
		       carpet_area_min, carpet_area_max, price_min, price_max,
		       COALESCE(ST_AsText(location), '') as location,
		       address, city, state, pincode, status,
		       amenities, media, created_at, updated_at
		FROM projects
		WHERE id = $1`

	p := &domain.Project{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.RERANumber, &p.BuilderID, &p.Description,
		&p.CarpetAreaMin, &p.CarpetAreaMax, &p.PriceMin, &p.PriceMax,
		&p.Location,
		&p.Address, &p.City, &p.State, &p.Pincode, &p.Status,
		&p.Amenities, &p.Media, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get project by id: %w", err)
	}

	return p, nil
}

// parseLocation parses a "lng,lat" string into two float64 values.
// Returns (0, 0, false) if the string is empty or invalid.
func parseLocation(loc string) (lng, lat float64, ok bool) {
	loc = strings.TrimSpace(loc)
	if loc == "" {
		return 0, 0, false
	}

	parts := strings.Split(loc, ",")
	if len(parts) != 2 {
		return 0, 0, false
	}

	lng, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, 0, false
	}

	lat, err = strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, 0, false
	}

	return lng, lat, true
}

// CreateProject inserts a new project into the projects table.
func (r *ProjectRepo) CreateProject(ctx context.Context, project *domain.Project) error {
	if project.Amenities == nil {
		project.Amenities = json.RawMessage("[]")
	}
	if project.Media == nil {
		project.Media = json.RawMessage("[]")
	}

	lng, lat, hasLocation := parseLocation(project.Location)

	var query string
	var args []interface{}

	if hasLocation {
		query = `
			INSERT INTO projects (
				id, name, slug, rera_number, builder_id, description,
				carpet_area_min, carpet_area_max, price_min, price_max,
				location, address, city, state, pincode, status,
				amenities, media, created_at, updated_at
			) VALUES (
				$1, $2, $3, $4, $5, $6,
				$7, $8, $9, $10,
				ST_SetSRID(ST_MakePoint($11, $12), 4326)::geography,
				$13, $14, $15, $16, $17,
				$18, $19, $20, $21
			)`
		args = []interface{}{
			project.ID, project.Name, project.Slug, project.RERANumber, project.BuilderID, project.Description,
			project.CarpetAreaMin, project.CarpetAreaMax, project.PriceMin, project.PriceMax,
			lng, lat,
			project.Address, project.City, project.State, project.Pincode, project.Status,
			project.Amenities, project.Media, project.CreatedAt, project.UpdatedAt,
		}
	} else {
		query = `
			INSERT INTO projects (
				id, name, slug, rera_number, builder_id, description,
				carpet_area_min, carpet_area_max, price_min, price_max,
				location, address, city, state, pincode, status,
				amenities, media, created_at, updated_at
			) VALUES (
				$1, $2, $3, $4, $5, $6,
				$7, $8, $9, $10,
				NULL,
				$11, $12, $13, $14, $15,
				$16, $17, $18, $19
			)`
		args = []interface{}{
			project.ID, project.Name, project.Slug, project.RERANumber, project.BuilderID, project.Description,
			project.CarpetAreaMin, project.CarpetAreaMax, project.PriceMin, project.PriceMax,
			project.Address, project.City, project.State, project.Pincode, project.Status,
			project.Amenities, project.Media, project.CreatedAt, project.UpdatedAt,
		}
	}

	_, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("create project: %w", err)
	}

	return nil
}

// UpdateProject updates all fields of an existing project by its ID.
func (r *ProjectRepo) UpdateProject(ctx context.Context, project *domain.Project) error {
	if project.Amenities == nil {
		project.Amenities = json.RawMessage("[]")
	}
	if project.Media == nil {
		project.Media = json.RawMessage("[]")
	}

	lng, lat, hasLocation := parseLocation(project.Location)

	var query string
	var args []interface{}

	if hasLocation {
		query = `
			UPDATE projects SET
				name = $1, slug = $2, rera_number = $3, builder_id = $4, description = $5,
				carpet_area_min = $6, carpet_area_max = $7, price_min = $8, price_max = $9,
				location = ST_SetSRID(ST_MakePoint($10, $11), 4326)::geography,
				address = $12, city = $13, state = $14, pincode = $15, status = $16,
				amenities = $17, media = $18, updated_at = $19
			WHERE id = $20`
		args = []interface{}{
			project.Name, project.Slug, project.RERANumber, project.BuilderID, project.Description,
			project.CarpetAreaMin, project.CarpetAreaMax, project.PriceMin, project.PriceMax,
			lng, lat,
			project.Address, project.City, project.State, project.Pincode, project.Status,
			project.Amenities, project.Media, project.UpdatedAt,
			project.ID,
		}
	} else {
		query = `
			UPDATE projects SET
				name = $1, slug = $2, rera_number = $3, builder_id = $4, description = $5,
				carpet_area_min = $6, carpet_area_max = $7, price_min = $8, price_max = $9,
				location = NULL,
				address = $10, city = $11, state = $12, pincode = $13, status = $14,
				amenities = $15, media = $16, updated_at = $17
			WHERE id = $18`
		args = []interface{}{
			project.Name, project.Slug, project.RERANumber, project.BuilderID, project.Description,
			project.CarpetAreaMin, project.CarpetAreaMax, project.PriceMin, project.PriceMax,
			project.Address, project.City, project.State, project.Pincode, project.Status,
			project.Amenities, project.Media, project.UpdatedAt,
			project.ID,
		}
	}

	_, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("update project: %w", err)
	}

	return nil
}

// GetProjectUnits returns all units belonging to a project.
func (r *ProjectRepo) GetProjectUnits(ctx context.Context, projectID uuid.UUID) ([]*domain.ProjectUnit, error) {
	query := `
		SELECT id, project_id, unit_number, unit_type, floor, carpet_area, price, status,
		       created_at, updated_at
		FROM project_units
		WHERE project_id = $1
		ORDER BY created_at ASC`

	rows, err := r.pool.Query(ctx, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("get project units: %w", err)
	}
	defer rows.Close()

	var units []*domain.ProjectUnit
	for rows.Next() {
		u := &domain.ProjectUnit{}
		err := rows.Scan(
			&u.ID, &u.ProjectID, &u.Name, &u.Type, &u.Floor,
			&u.CarpetArea, &u.Price, &u.Status,
			&u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("get project units scan: %w", err)
		}
		// Amenities column does not exist in project_units table; default to empty array.
		u.Amenities = json.RawMessage("[]")
		units = append(units, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get project units rows: %w", err)
	}

	return units, nil
}

// GetDueDiligence returns the due-diligence report for a project.
// Currently there is no due_diligence table, so this always returns a not-found error.
func (r *ProjectRepo) GetDueDiligence(ctx context.Context, projectID uuid.UUID) (*domain.DueDiligenceReport, error) {
	return nil, fmt.Errorf("due diligence report not found for project %s", projectID)
}

// Compile-time check that ProjectRepo satisfies service.ProjectRepository.
var _ service.ProjectRepository = (*ProjectRepo)(nil)
