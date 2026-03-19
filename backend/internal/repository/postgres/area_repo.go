package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/service"
)

// AreaRepo implements service.AreaRepository using raw SQL against a pgx pool.
type AreaRepo struct {
	pool *pgxpool.Pool
}

// NewAreaRepo creates a new AreaRepo backed by the given connection pool.
func NewAreaRepo(pool *pgxpool.Pool) *AreaRepo {
	return &AreaRepo{pool: pool}
}

// CreateArea inserts a new area into the areas table.
func (r *AreaRepo) CreateArea(ctx context.Context, area *domain.Area) error {
	if area.PriceTrend == nil {
		area.PriceTrend = []byte("[]")
	}
	if area.Infrastructure == nil {
		area.Infrastructure = []byte("[]")
	}

	lng, lat, hasLocation := parseLocation(area.Location)

	var query string
	var args []interface{}

	if hasLocation {
		query = `
			INSERT INTO areas (
				id, name, slug, city, state, description,
				location, price_trend, infrastructure,
				created_at, updated_at
			) VALUES (
				$1, $2, $3, $4, $5, $6,
				ST_SetSRID(ST_MakePoint($7, $8), 4326)::geography,
				$9, $10, $11, $12
			)`
		args = []interface{}{
			area.ID, area.Name, area.Slug, area.City, area.State, area.Description,
			lng, lat,
			area.PriceTrend, area.Infrastructure,
			area.CreatedAt, area.UpdatedAt,
		}
	} else {
		query = `
			INSERT INTO areas (
				id, name, slug, city, state, description,
				location, price_trend, infrastructure,
				created_at, updated_at
			) VALUES (
				$1, $2, $3, $4, $5, $6,
				NULL, $7, $8, $9, $10
			)`
		args = []interface{}{
			area.ID, area.Name, area.Slug, area.City, area.State, area.Description,
			area.PriceTrend, area.Infrastructure,
			area.CreatedAt, area.UpdatedAt,
		}
	}

	_, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("create area: %w", err)
	}

	return nil
}

// GetAreaBySlug retrieves a single area by its URL-friendly slug.
func (r *AreaRepo) GetAreaBySlug(ctx context.Context, slug string) (*domain.Area, error) {
	query := `
		SELECT id, name, slug, city, state, description,
		       COALESCE(ST_AsText(location), '') as location,
		       price_trend, infrastructure, created_at, updated_at
		FROM areas
		WHERE slug = $1`

	a := &domain.Area{}
	row := r.pool.QueryRow(ctx, query, slug)

	err := row.Scan(
		&a.ID, &a.Name, &a.Slug, &a.City, &a.State, &a.Description,
		&a.Location,
		&a.PriceTrend, &a.Infrastructure, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get area by slug: %w", err)
	}

	return a, nil
}

// ListAreas returns a filtered, paginated list of areas along with the total count.
func (r *AreaRepo) ListAreas(ctx context.Context, filters service.AreaFilters) ([]*domain.Area, int, error) {
	query := `
		SELECT id, name, slug, city, state, description,
		       COALESCE(ST_AsText(location), '') as location,
		       price_trend, infrastructure, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM areas`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.City != "" {
		conditions = append(conditions, fmt.Sprintf("city ILIKE $%d", argIdx))
		args = append(args, "%"+filters.City+"%")
		argIdx++
	}
	if filters.State != "" {
		conditions = append(conditions, fmt.Sprintf("state ILIKE $%d", argIdx))
		args = append(args, "%"+filters.State+"%")
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
		return nil, 0, fmt.Errorf("list areas: %w", err)
	}
	defer rows.Close()

	var areas []*domain.Area
	var total int

	for rows.Next() {
		a := &domain.Area{}
		err := rows.Scan(
			&a.ID, &a.Name, &a.Slug, &a.City, &a.State, &a.Description,
			&a.Location,
			&a.PriceTrend, &a.Infrastructure, &a.CreatedAt, &a.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list areas scan: %w", err)
		}
		areas = append(areas, a)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list areas rows: %w", err)
	}

	return areas, total, nil
}

// Compile-time check that AreaRepo satisfies service.AreaRepository.
var _ service.AreaRepository = (*AreaRepo)(nil)
