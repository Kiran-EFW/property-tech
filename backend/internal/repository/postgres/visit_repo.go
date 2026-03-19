package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/service"
)

// VisitRepo implements service.VisitRepository using raw SQL against a pgx pool.
type VisitRepo struct {
	pool *pgxpool.Pool
}

// NewVisitRepo creates a new VisitRepo backed by the given connection pool.
func NewVisitRepo(pool *pgxpool.Pool) *VisitRepo {
	return &VisitRepo{pool: pool}
}

// CreateVisit inserts a new site visit into the site_visits table.
func (r *VisitRepo) CreateVisit(ctx context.Context, visit *domain.SiteVisit) error {
	photos := visit.Photos
	if photos == nil {
		photos = json.RawMessage("[]")
	}

	query := `
		INSERT INTO site_visits (
			id, lead_id, agent_id, project_id, feedback,
			photos, duration, rating, visited_at, created_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10
		)`

	_, err := r.pool.Exec(ctx, query,
		visit.ID,
		visit.LeadID,
		visit.AgentID,
		visit.ProjectID,
		visit.Feedback,
		photos,
		visit.Duration,
		visit.Rating,
		visit.VisitedAt,
		visit.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("create visit: %w", err)
	}

	return nil
}

// GetVisitByID retrieves a single site visit by its unique identifier.
func (r *VisitRepo) GetVisitByID(ctx context.Context, id uuid.UUID) (*domain.SiteVisit, error) {
	query := `
		SELECT id, lead_id, agent_id, project_id, feedback,
		       photos, duration, rating, visited_at, created_at
		FROM site_visits
		WHERE id = $1`

	v := &domain.SiteVisit{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&v.ID,
		&v.LeadID,
		&v.AgentID,
		&v.ProjectID,
		&v.Feedback,
		&v.Photos,
		&v.Duration,
		&v.Rating,
		&v.VisitedAt,
		&v.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get visit by id: %w", err)
	}

	return v, nil
}

// ListVisits returns a filtered, paginated list of site visits along with the total count.
func (r *VisitRepo) ListVisits(ctx context.Context, filters service.VisitFilters) ([]*domain.SiteVisit, int, error) {
	query := `
		SELECT id, lead_id, agent_id, project_id, feedback,
		       photos, duration, rating, visited_at, created_at,
		       COUNT(*) OVER() as total_count
		FROM site_visits`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.ProjectID != "" {
		pid, err := uuid.Parse(filters.ProjectID)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits: invalid project_id filter: %w", err)
		}
		conditions = append(conditions, fmt.Sprintf("project_id = $%d", argIdx))
		args = append(args, pid)
		argIdx++
	}
	if filters.LeadID != "" {
		lid, err := uuid.Parse(filters.LeadID)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits: invalid lead_id filter: %w", err)
		}
		conditions = append(conditions, fmt.Sprintf("lead_id = $%d", argIdx))
		args = append(args, lid)
		argIdx++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY visited_at DESC"

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
		return nil, 0, fmt.Errorf("list visits: %w", err)
	}
	defer rows.Close()

	var visits []*domain.SiteVisit
	var total int

	for rows.Next() {
		v := &domain.SiteVisit{}
		err := rows.Scan(
			&v.ID,
			&v.LeadID,
			&v.AgentID,
			&v.ProjectID,
			&v.Feedback,
			&v.Photos,
			&v.Duration,
			&v.Rating,
			&v.VisitedAt,
			&v.CreatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits scan: %w", err)
		}
		visits = append(visits, v)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list visits rows: %w", err)
	}

	return visits, total, nil
}

// ListVisitsByAgent returns site visits for a specific agent, with optional filters.
func (r *VisitRepo) ListVisitsByAgent(ctx context.Context, agentID uuid.UUID, filters service.VisitFilters) ([]*domain.SiteVisit, int, error) {
	query := `
		SELECT id, lead_id, agent_id, project_id, feedback,
		       photos, duration, rating, visited_at, created_at,
		       COUNT(*) OVER() as total_count
		FROM site_visits
		WHERE agent_id = $1`

	args := []interface{}{agentID}
	argIdx := 2

	if filters.ProjectID != "" {
		pid, err := uuid.Parse(filters.ProjectID)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits by agent: invalid project_id filter: %w", err)
		}
		query += fmt.Sprintf(" AND project_id = $%d", argIdx)
		args = append(args, pid)
		argIdx++
	}
	if filters.LeadID != "" {
		lid, err := uuid.Parse(filters.LeadID)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits by agent: invalid lead_id filter: %w", err)
		}
		query += fmt.Sprintf(" AND lead_id = $%d", argIdx)
		args = append(args, lid)
		argIdx++
	}

	query += " ORDER BY visited_at DESC"

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
		return nil, 0, fmt.Errorf("list visits by agent: %w", err)
	}
	defer rows.Close()

	var visits []*domain.SiteVisit
	var total int

	for rows.Next() {
		v := &domain.SiteVisit{}
		err := rows.Scan(
			&v.ID,
			&v.LeadID,
			&v.AgentID,
			&v.ProjectID,
			&v.Feedback,
			&v.Photos,
			&v.Duration,
			&v.Rating,
			&v.VisitedAt,
			&v.CreatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list visits by agent scan: %w", err)
		}
		visits = append(visits, v)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list visits by agent rows: %w", err)
	}

	return visits, total, nil
}

// UpdateVisit updates the feedback, photos, duration, and rating of an existing site visit.
func (r *VisitRepo) UpdateVisit(ctx context.Context, visit *domain.SiteVisit) error {
	query := `
		UPDATE site_visits SET
			feedback = $1, photos = $2, duration = $3, rating = $4
		WHERE id = $5`

	_, err := r.pool.Exec(ctx, query,
		visit.Feedback,
		visit.Photos,
		visit.Duration,
		visit.Rating,
		visit.ID,
	)
	if err != nil {
		return fmt.Errorf("update visit: %w", err)
	}

	return nil
}

// Compile-time check that VisitRepo satisfies service.VisitRepository.
var _ service.VisitRepository = (*VisitRepo)(nil)
