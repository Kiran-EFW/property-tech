package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/service"
)

// LeadRepo implements service.LeadRepository using raw SQL against a pgx pool.
type LeadRepo struct {
	pool *pgxpool.Pool
}

// NewLeadRepo creates a new LeadRepo backed by the given connection pool.
func NewLeadRepo(pool *pgxpool.Pool) *LeadRepo {
	return &LeadRepo{pool: pool}
}

// CreateLead inserts a new lead into the leads table.
func (r *LeadRepo) CreateLead(ctx context.Context, lead *domain.Lead) error {
	query := `
		INSERT INTO leads (
			id, investor_id, project_id, agent_id, source, status,
			phone, name, email, budget, notes, follow_up_at,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, NULL,
			$12, $13
		)`

	_, err := r.pool.Exec(ctx, query,
		lead.ID,
		lead.InvestorID,
		lead.ProjectID,
		lead.AgentID,
		lead.Source,
		lead.Status,
		lead.Phone,
		lead.Name,
		lead.Email,
		lead.Budget,
		lead.Notes,
		lead.CreatedAt,
		lead.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create lead: %w", err)
	}

	return nil
}

// GetLeadByID retrieves a single lead by its unique identifier.
func (r *LeadRepo) GetLeadByID(ctx context.Context, id uuid.UUID) (*domain.Lead, error) {
	query := `
		SELECT id, investor_id, project_id, agent_id, source, status,
		       phone, name, email, budget, notes, created_at, updated_at
		FROM leads
		WHERE id = $1`

	l := &domain.Lead{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&l.ID,
		&l.InvestorID,
		&l.ProjectID,
		&l.AgentID,
		&l.Source,
		&l.Status,
		&l.Phone,
		&l.Name,
		&l.Email,
		&l.Budget,
		&l.Notes,
		&l.CreatedAt,
		&l.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get lead by id: %w", err)
	}

	return l, nil
}

// ListLeads returns a filtered, paginated list of leads along with the total count.
func (r *LeadRepo) ListLeads(ctx context.Context, filters service.LeadFilters) ([]*domain.Lead, int, error) {
	query := `
		SELECT id, investor_id, project_id, agent_id, source, status,
		       phone, name, email, budget, notes, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM leads`

	conditions := []string{}
	args := []interface{}{}
	argIdx := 1

	if filters.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, filters.Status)
		argIdx++
	}
	if filters.ProjectID != "" {
		pid, err := uuid.Parse(filters.ProjectID)
		if err != nil {
			return nil, 0, fmt.Errorf("list leads: invalid project_id filter: %w", err)
		}
		conditions = append(conditions, fmt.Sprintf("project_id = $%d", argIdx))
		args = append(args, pid)
		argIdx++
	}
	if filters.AgentID != "" {
		aid, err := uuid.Parse(filters.AgentID)
		if err != nil {
			return nil, 0, fmt.Errorf("list leads: invalid agent_id filter: %w", err)
		}
		conditions = append(conditions, fmt.Sprintf("agent_id = $%d", argIdx))
		args = append(args, aid)
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
		return nil, 0, fmt.Errorf("list leads: %w", err)
	}
	defer rows.Close()

	var leads []*domain.Lead
	var total int

	for rows.Next() {
		l := &domain.Lead{}
		err := rows.Scan(
			&l.ID,
			&l.InvestorID,
			&l.ProjectID,
			&l.AgentID,
			&l.Source,
			&l.Status,
			&l.Phone,
			&l.Name,
			&l.Email,
			&l.Budget,
			&l.Notes,
			&l.CreatedAt,
			&l.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list leads scan: %w", err)
		}
		leads = append(leads, l)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list leads rows: %w", err)
	}

	return leads, total, nil
}

// ListLeadsByAgent returns leads assigned to a specific agent, with optional status filter.
func (r *LeadRepo) ListLeadsByAgent(ctx context.Context, agentID uuid.UUID, filters service.LeadFilters) ([]*domain.Lead, int, error) {
	query := `
		SELECT id, investor_id, project_id, agent_id, source, status,
		       phone, name, email, budget, notes, created_at, updated_at,
		       COUNT(*) OVER() as total_count
		FROM leads
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
		return nil, 0, fmt.Errorf("list leads by agent: %w", err)
	}
	defer rows.Close()

	var leads []*domain.Lead
	var total int

	for rows.Next() {
		l := &domain.Lead{}
		err := rows.Scan(
			&l.ID,
			&l.InvestorID,
			&l.ProjectID,
			&l.AgentID,
			&l.Source,
			&l.Status,
			&l.Phone,
			&l.Name,
			&l.Email,
			&l.Budget,
			&l.Notes,
			&l.CreatedAt,
			&l.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("list leads by agent scan: %w", err)
		}
		leads = append(leads, l)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("list leads by agent rows: %w", err)
	}

	return leads, total, nil
}

// UpdateLead updates the mutable fields of an existing lead.
func (r *LeadRepo) UpdateLead(ctx context.Context, lead *domain.Lead) error {
	query := `
		UPDATE leads SET
			investor_id = $1, agent_id = $2, status = $3,
			phone = $4, name = $5, email = $6,
			budget = $7, notes = $8, updated_at = $9
		WHERE id = $10`

	_, err := r.pool.Exec(ctx, query,
		lead.InvestorID,
		lead.AgentID,
		lead.Status,
		lead.Phone,
		lead.Name,
		lead.Email,
		lead.Budget,
		lead.Notes,
		lead.UpdatedAt,
		lead.ID,
	)
	if err != nil {
		return fmt.Errorf("update lead: %w", err)
	}

	return nil
}

// AddLeadNote inserts a new note for a lead.
func (r *LeadRepo) AddLeadNote(ctx context.Context, note *domain.LeadNote) error {
	query := `
		INSERT INTO lead_notes (id, lead_id, author_id, content, created_at)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := r.pool.Exec(ctx, query,
		note.ID,
		note.LeadID,
		note.AuthorID,
		note.Content,
		note.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("add lead note: %w", err)
	}

	return nil
}

// GetLeadByPhoneAndProject retrieves a lead by phone number and project ID.
func (r *LeadRepo) GetLeadByPhoneAndProject(ctx context.Context, phone string, projectID uuid.UUID) (*domain.Lead, error) {
	query := `
		SELECT id, investor_id, project_id, agent_id, source, status,
		       phone, name, email, budget, notes, created_at, updated_at
		FROM leads
		WHERE phone = $1 AND project_id = $2
		LIMIT 1`

	l := &domain.Lead{}
	row := r.pool.QueryRow(ctx, query, phone, projectID)

	err := row.Scan(
		&l.ID,
		&l.InvestorID,
		&l.ProjectID,
		&l.AgentID,
		&l.Source,
		&l.Status,
		&l.Phone,
		&l.Name,
		&l.Email,
		&l.Budget,
		&l.Notes,
		&l.CreatedAt,
		&l.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get lead by phone and project: %w", err)
	}

	return l, nil
}

// GetAgentWithFewestActiveLeads finds the active agent with the fewest leads that are
// not in a terminal state (converted or lost).
func (r *LeadRepo) GetAgentWithFewestActiveLeads(ctx context.Context) (*domain.Agent, error) {
	query := `
		SELECT a.id, a.user_id, a.rera_number, a.pan, a.gst,
		       a.tier, a.is_active, a.created_at, a.updated_at
		FROM agents a
		LEFT JOIN leads l ON l.agent_id = a.id
			AND l.status NOT IN ('converted', 'lost')
		WHERE a.is_active = true
		GROUP BY a.id, a.user_id, a.rera_number, a.pan, a.gst,
		         a.tier, a.is_active, a.created_at, a.updated_at
		ORDER BY COUNT(l.id) ASC
		LIMIT 1`

	agent := &domain.Agent{}
	row := r.pool.QueryRow(ctx, query)

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
		return nil, fmt.Errorf("get agent with fewest active leads: %w", err)
	}

	return agent, nil
}

// GetUncontactedLeadsOlderThan returns leads in "new" status that were created
// before (now - threshold). Used by the lead escalation worker.
func (r *LeadRepo) GetUncontactedLeadsOlderThan(ctx context.Context, threshold time.Duration) ([]*domain.Lead, error) {
	query := `
		SELECT id, investor_id, project_id, agent_id, source, status,
		       phone, name, email, budget, notes, created_at, updated_at
		FROM leads
		WHERE status = 'new'
		  AND created_at < now() - $1::interval
		ORDER BY created_at ASC`

	rows, err := r.pool.Query(ctx, query, threshold.String())
	if err != nil {
		return nil, fmt.Errorf("get uncontacted leads: %w", err)
	}
	defer rows.Close()

	var leads []*domain.Lead
	for rows.Next() {
		l := &domain.Lead{}
		err := rows.Scan(
			&l.ID,
			&l.InvestorID,
			&l.ProjectID,
			&l.AgentID,
			&l.Source,
			&l.Status,
			&l.Phone,
			&l.Name,
			&l.Email,
			&l.Budget,
			&l.Notes,
			&l.CreatedAt,
			&l.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("get uncontacted leads scan: %w", err)
		}
		leads = append(leads, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get uncontacted leads rows: %w", err)
	}

	return leads, nil
}

// Compile-time check that LeadRepo satisfies service.LeadRepository.
var _ service.LeadRepository = (*LeadRepo)(nil)
