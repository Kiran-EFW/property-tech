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

// AgentRepo implements service.AgentRepository using raw SQL against a pgx pool.
type AgentRepo struct {
	pool *pgxpool.Pool
}

// NewAgentRepo creates a new AgentRepo backed by the given connection pool.
func NewAgentRepo(pool *pgxpool.Pool) *AgentRepo {
	return &AgentRepo{pool: pool}
}

// CreateAgent inserts a new agent into the agents table.
func (r *AgentRepo) CreateAgent(ctx context.Context, agent *domain.Agent) error {
	query := `
		INSERT INTO agents (
			id, user_id, rera_number, pan, gst,
			tier, is_active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9
		)`

	_, err := r.pool.Exec(ctx, query,
		agent.ID,
		agent.UserID,
		agent.RERANumber,
		agent.PAN,
		agent.GST,
		agent.Tier,
		agent.IsActive,
		agent.CreatedAt,
		agent.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create agent: %w", err)
	}

	return nil
}

// GetAgentByID retrieves a single agent by its unique identifier.
func (r *AgentRepo) GetAgentByID(ctx context.Context, id uuid.UUID) (*domain.Agent, error) {
	query := `
		SELECT id, user_id, rera_number, pan, gst,
		       tier, is_active, created_at, updated_at
		FROM agents
		WHERE id = $1`

	agent := &domain.Agent{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&agent.ID,
		&agent.UserID,
		&agent.RERANumber,
		&agent.PAN,
		&agent.GST,
		&agent.Tier,
		&agent.IsActive,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get agent by id: %w", err)
	}

	return agent, nil
}

// GetAgentByUserID retrieves an agent by the associated user ID.
func (r *AgentRepo) GetAgentByUserID(ctx context.Context, userID uuid.UUID) (*domain.Agent, error) {
	query := `
		SELECT id, user_id, rera_number, pan, gst,
		       tier, is_active, created_at, updated_at
		FROM agents
		WHERE user_id = $1`

	agent := &domain.Agent{}
	row := r.pool.QueryRow(ctx, query, userID)

	err := row.Scan(
		&agent.ID,
		&agent.UserID,
		&agent.RERANumber,
		&agent.PAN,
		&agent.GST,
		&agent.Tier,
		&agent.IsActive,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get agent by user id: %w", err)
	}

	return agent, nil
}

// ListAgents returns a filtered, paginated list of agents along with the total count.
func (r *AgentRepo) ListAgents(ctx context.Context, filters service.AgentFilters) ([]*domain.Agent, int, error) {
	query := `
		SELECT id, user_id, rera_number, pan, gst,
		       tier, is_active, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM agents`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.Tier != "" {
		conditions = append(conditions, fmt.Sprintf("tier = $%d", argIdx))
		args = append(args, filters.Tier)
		argIdx++
	}
	if filters.IsActive != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", argIdx))
		args = append(args, *filters.IsActive)
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
		return nil, 0, fmt.Errorf("list agents: %w", err)
	}
	defer rows.Close()

	var agents []*domain.Agent
	var total int

	for rows.Next() {
		a := &domain.Agent{}
		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.RERANumber,
			&a.PAN,
			&a.GST,
			&a.Tier,
			&a.IsActive,
			&a.CreatedAt,
			&a.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list agents scan: %w", err)
		}
		agents = append(agents, a)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list agents rows: %w", err)
	}

	return agents, total, nil
}

// UpdateAgent updates all mutable fields of an existing agent.
func (r *AgentRepo) UpdateAgent(ctx context.Context, agent *domain.Agent) error {
	query := `
		UPDATE agents SET
			rera_number = $1, pan = $2, gst = $3,
			tier = $4, is_active = $5, updated_at = $6
		WHERE id = $7`

	_, err := r.pool.Exec(ctx, query,
		agent.RERANumber,
		agent.PAN,
		agent.GST,
		agent.Tier,
		agent.IsActive,
		agent.UpdatedAt,
		agent.ID,
	)
	if err != nil {
		return fmt.Errorf("update agent: %w", err)
	}

	return nil
}

// parsePeriodDays converts a period string like "30d", "90d", or "1y" into a number of days.
func parsePeriodDays(period string) (int, error) {
	if period == "" {
		return 30, nil
	}

	period = strings.TrimSpace(period)
	if strings.HasSuffix(period, "y") {
		numStr := strings.TrimSuffix(period, "y")
		var years int
		if _, err := fmt.Sscanf(numStr, "%d", &years); err != nil {
			return 0, fmt.Errorf("invalid period %q: %w", period, err)
		}
		return years * 365, nil
	}

	if strings.HasSuffix(period, "d") {
		numStr := strings.TrimSuffix(period, "d")
		var days int
		if _, err := fmt.Sscanf(numStr, "%d", &days); err != nil {
			return 0, fmt.Errorf("invalid period %q: %w", period, err)
		}
		return days, nil
	}

	return 0, fmt.Errorf("invalid period format %q: expected suffix 'd' or 'y'", period)
}

// GetAgentPerformance computes performance metrics for an agent over the given period.
func (r *AgentRepo) GetAgentPerformance(ctx context.Context, agentID uuid.UUID, period string) (*service.AgentPerformance, error) {
	days, err := parsePeriodDays(period)
	if err != nil {
		return nil, fmt.Errorf("get agent performance: %w", err)
	}

	query := `
		WITH period_leads AS (
			SELECT id, status
			FROM leads
			WHERE agent_id = $1
			  AND created_at >= NOW() - ($2 || ' days')::interval
		),
		lead_counts AS (
			SELECT
				COUNT(*) AS total_leads,
				COUNT(*) FILTER (WHERE status NOT IN ('converted', 'lost')) AS active_leads,
				COUNT(*) FILTER (WHERE status = 'site_visit') AS site_visits,
				COUNT(*) FILTER (WHERE status = 'converted') AS bookings
			FROM period_leads
		),
		earnings AS (
			SELECT COALESCE(SUM(c.amount), 0) AS total_earnings
			FROM commissions c
			JOIN bookings b ON b.id = c.booking_id
			WHERE c.agent_id = $1
			  AND c.created_at >= NOW() - ($2 || ' days')::interval
		)
		SELECT
			lc.total_leads,
			lc.active_leads,
			lc.site_visits,
			lc.bookings,
			e.total_earnings
		FROM lead_counts lc, earnings e`

	perf := &service.AgentPerformance{
		AgentID: agentID,
		Period:  period,
	}

	row := r.pool.QueryRow(ctx, query, agentID, days)
	err = row.Scan(
		&perf.TotalLeads,
		&perf.ActiveLeads,
		&perf.SiteVisits,
		&perf.Bookings,
		&perf.TotalEarnings,
	)
	if err != nil {
		return nil, fmt.Errorf("get agent performance: %w", err)
	}

	if perf.TotalLeads > 0 {
		perf.ConversionRate = float64(perf.Bookings) / float64(perf.TotalLeads)
	}

	return perf, nil
}

// Compile-time check that AgentRepo satisfies service.AgentRepository.
var _ service.AgentRepository = (*AgentRepo)(nil)
