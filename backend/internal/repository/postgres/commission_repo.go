package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/service"
)

// CommissionRepo implements service.CommissionRepository using raw SQL against a pgx pool.
type CommissionRepo struct {
	pool *pgxpool.Pool
}

// NewCommissionRepo creates a new CommissionRepo backed by the given connection pool.
func NewCommissionRepo(pool *pgxpool.Pool) *CommissionRepo {
	return &CommissionRepo{pool: pool}
}

// CreateCommission inserts a new commission into the commissions table.
func (r *CommissionRepo) CreateCommission(ctx context.Context, commission *domain.Commission) error {
	query := `
		INSERT INTO commissions (
			id, booking_id, agent_id, amount, tds,
			net_amount, status, paid_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10
		)`

	_, err := r.pool.Exec(ctx, query,
		commission.ID,
		commission.BookingID,
		commission.AgentID,
		commission.Amount,
		commission.TDS,
		commission.NetAmount,
		commission.Status,
		commission.PaidAt,
		commission.CreatedAt,
		commission.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create commission: %w", err)
	}

	return nil
}

// GetCommissionByID retrieves a single commission by its unique identifier.
func (r *CommissionRepo) GetCommissionByID(ctx context.Context, id uuid.UUID) (*domain.Commission, error) {
	query := `
		SELECT id, booking_id, agent_id, amount, tds,
		       net_amount, status, paid_at, created_at, updated_at
		FROM commissions
		WHERE id = $1`

	c := &domain.Commission{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&c.ID,
		&c.BookingID,
		&c.AgentID,
		&c.Amount,
		&c.TDS,
		&c.NetAmount,
		&c.Status,
		&c.PaidAt,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get commission by id: %w", err)
	}

	return c, nil
}

// ListCommissions returns a filtered, paginated list of commissions along with the total count.
func (r *CommissionRepo) ListCommissions(ctx context.Context, filters service.CommissionFilters) ([]*domain.Commission, int, error) {
	query := `
		SELECT id, booking_id, agent_id, amount, tds,
		       net_amount, status, paid_at, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM commissions`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.AgentID != "" {
		aid, err := uuid.Parse(filters.AgentID)
		if err != nil {
			return nil, 0, fmt.Errorf("list commissions: invalid agent_id filter: %w", err)
		}
		conditions = append(conditions, fmt.Sprintf("agent_id = $%d", argIdx))
		args = append(args, aid)
		argIdx++
	}
	if filters.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, filters.Status)
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
		return nil, 0, fmt.Errorf("list commissions: %w", err)
	}
	defer rows.Close()

	var commissions []*domain.Commission
	var total int

	for rows.Next() {
		c := &domain.Commission{}
		err := rows.Scan(
			&c.ID,
			&c.BookingID,
			&c.AgentID,
			&c.Amount,
			&c.TDS,
			&c.NetAmount,
			&c.Status,
			&c.PaidAt,
			&c.CreatedAt,
			&c.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list commissions scan: %w", err)
		}
		commissions = append(commissions, c)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list commissions rows: %w", err)
	}

	return commissions, total, nil
}

// ListCommissionsByAgent returns commissions for a specific agent, with optional status filter.
func (r *CommissionRepo) ListCommissionsByAgent(ctx context.Context, agentID uuid.UUID, filters service.CommissionFilters) ([]*domain.Commission, int, error) {
	query := `
		SELECT id, booking_id, agent_id, amount, tds,
		       net_amount, status, paid_at, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM commissions
		WHERE agent_id = $1`

	args := []interface{}{agentID}
	argIdx := 2

	if filters.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, filters.Status)
		argIdx++
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
		return nil, 0, fmt.Errorf("list commissions by agent: %w", err)
	}
	defer rows.Close()

	var commissions []*domain.Commission
	var total int

	for rows.Next() {
		c := &domain.Commission{}
		err := rows.Scan(
			&c.ID,
			&c.BookingID,
			&c.AgentID,
			&c.Amount,
			&c.TDS,
			&c.NetAmount,
			&c.Status,
			&c.PaidAt,
			&c.CreatedAt,
			&c.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list commissions by agent scan: %w", err)
		}
		commissions = append(commissions, c)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list commissions by agent rows: %w", err)
	}

	return commissions, total, nil
}

// UpdateCommission updates the status, paid_at, and updated_at fields of an existing commission.
func (r *CommissionRepo) UpdateCommission(ctx context.Context, commission *domain.Commission) error {
	query := `
		UPDATE commissions SET
			status = $1, paid_at = $2, updated_at = $3
		WHERE id = $4`

	_, err := r.pool.Exec(ctx, query,
		commission.Status,
		commission.PaidAt,
		commission.UpdatedAt,
		commission.ID,
	)
	if err != nil {
		return fmt.Errorf("update commission: %w", err)
	}

	return nil
}

// GetCommissionSummary returns aggregated commission statistics.
// If agentID is non-nil, results are scoped to that agent.
func (r *CommissionRepo) GetCommissionSummary(ctx context.Context, agentID *uuid.UUID) (*service.CommissionSummary, error) {
	query := `
		SELECT
			COALESCE(SUM(amount), 0) AS total_earned,
			COALESCE(SUM(amount) FILTER (WHERE status = 'pending'), 0) AS total_pending,
			COALESCE(SUM(amount) FILTER (WHERE status = 'approved'), 0) AS total_approved,
			COALESCE(SUM(amount) FILTER (WHERE status = 'paid'), 0) AS total_paid,
			COUNT(*) AS count
		FROM commissions`

	args := []interface{}{}

	if agentID != nil {
		query += " WHERE agent_id = $1"
		args = append(args, *agentID)
	}

	summary := &service.CommissionSummary{}
	row := r.pool.QueryRow(ctx, query, args...)

	err := row.Scan(
		&summary.TotalEarned,
		&summary.TotalPending,
		&summary.TotalApproved,
		&summary.TotalPaid,
		&summary.Count,
	)
	if err != nil {
		return nil, fmt.Errorf("get commission summary: %w", err)
	}

	return summary, nil
}

// Compile-time check that CommissionRepo satisfies service.CommissionRepository.
var _ service.CommissionRepository = (*CommissionRepo)(nil)
