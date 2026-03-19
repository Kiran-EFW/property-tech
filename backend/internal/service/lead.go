package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// LeadFilters holds the query filters for listing leads.
type LeadFilters struct {
	Status    string `query:"status"`
	ProjectID string `query:"project_id"`
	AgentID   string `query:"agent_id"`
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
}

// CreateLeadInput holds the fields needed to create a lead.
type CreateLeadInput struct {
	ProjectID string  `json:"project_id" validate:"required,uuid"`
	Phone     string  `json:"phone" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Email     *string `json:"email"`
	Budget    *float64 `json:"budget"`
	Source    string  `json:"source"`
	Notes     *string `json:"notes"`
}

// UpdateLeadStatusInput holds the status transition payload.
type UpdateLeadStatusInput struct {
	Status  string `json:"status" validate:"required"`
	Remarks string `json:"remarks"`
}

// AddNoteInput holds the payload for adding a note to a lead.
type AddNoteInput struct {
	Content string `json:"content" validate:"required"`
}

// AssignLeadInput holds the payload for assigning a lead to an agent.
type AssignLeadInput struct {
	AgentID string `json:"agent_id" validate:"required,uuid"`
}

// LeadRepository defines the database operations required by LeadService.
type LeadRepository interface {
	CreateLead(ctx context.Context, lead *domain.Lead) error
	GetLeadByID(ctx context.Context, id uuid.UUID) (*domain.Lead, error)
	ListLeads(ctx context.Context, filters LeadFilters) ([]*domain.Lead, int, error)
	ListLeadsByAgent(ctx context.Context, agentID uuid.UUID, filters LeadFilters) ([]*domain.Lead, int, error)
	UpdateLead(ctx context.Context, lead *domain.Lead) error
	AddLeadNote(ctx context.Context, note *domain.LeadNote) error
	GetLeadByPhoneAndProject(ctx context.Context, phone string, projectID uuid.UUID) (*domain.Lead, error)
	GetAgentWithFewestActiveLeads(ctx context.Context) (*domain.Agent, error)
}

// LeadService handles lead management business logic.
type LeadService struct {
	repo     LeadRepository
	eventSvc *EventService
}

// NewLeadService creates a new LeadService.
func NewLeadService(repo LeadRepository, eventSvc *EventService) *LeadService {
	return &LeadService{
		repo:     repo,
		eventSvc: eventSvc,
	}
}

// Create creates a new lead, normalizes the phone number, deduplicates, and auto-assigns an agent.
func (s *LeadService) Create(ctx context.Context, input CreateLeadInput) (*domain.Lead, error) {
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("invalid project_id: %w", err)
	}

	// Normalize phone number: ensure it has country code prefix.
	phone := normalizePhone(input.Phone)

	// Check for duplicate lead (same phone + project).
	existing, err := s.repo.GetLeadByPhoneAndProject(ctx, phone, projectID)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("lead already exists for this phone number and project")
	}

	source := domain.LeadSource(input.Source)
	if source == "" {
		source = domain.LeadSourceWebsite
	}

	lead := &domain.Lead{
		ID:        uuid.New(),
		ProjectID: projectID,
		Phone:     phone,
		Name:      input.Name,
		Email:     input.Email,
		Budget:    input.Budget,
		Source:    source,
		Notes:     input.Notes,
		Status:    domain.LeadStatusNew,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Auto-assign agent: find agent with fewest active leads.
	agent, err := s.repo.GetAgentWithFewestActiveLeads(ctx)
	if err == nil && agent != nil {
		lead.AgentID = &agent.ID
		log.Info().Str("agent_id", agent.ID.String()).Str("lead_id", lead.ID.String()).Msg("lead auto-assigned to agent")
	}

	if err := s.repo.CreateLead(ctx, lead); err != nil {
		return nil, fmt.Errorf("failed to create lead: %w", err)
	}

	log.Info().Str("lead_id", lead.ID.String()).Str("phone", phone).Msg("lead created")
	return lead, nil
}

// List returns leads scoped by role: agents see only their leads, admins see all.
func (s *LeadService) List(ctx context.Context, filters LeadFilters, role string, userID uuid.UUID) ([]*domain.Lead, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	if role == "admin" || role == "super_admin" {
		return s.repo.ListLeads(ctx, filters)
	}

	// Agent sees only their assigned leads.
	return s.repo.ListLeadsByAgent(ctx, userID, filters)
}

// GetByID returns a lead by ID.
func (s *LeadService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Lead, error) {
	lead, err := s.repo.GetLeadByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("lead not found: %w", err)
	}
	return lead, nil
}

// UpdateStatus transitions a lead to a new status and creates an event log entry.
func (s *LeadService) UpdateStatus(ctx context.Context, id uuid.UUID, status, remarks string, actorID uuid.UUID) error {
	lead, err := s.repo.GetLeadByID(ctx, id)
	if err != nil {
		return fmt.Errorf("lead not found: %w", err)
	}

	lead.Status = domain.LeadStatus(status)
	lead.UpdatedAt = time.Now()

	if err := s.repo.UpdateLead(ctx, lead); err != nil {
		return fmt.Errorf("failed to update lead status: %w", err)
	}

	// Log the status change event.
	if s.eventSvc != nil {
		_ = s.eventSvc.Log(ctx, actorID, "agent", "lead_status_updated", "lead", id, map[string]interface{}{
			"new_status": status,
			"remarks":    remarks,
		})
	}

	log.Info().Str("lead_id", id.String()).Str("status", status).Msg("lead status updated")
	return nil
}

// AddNote adds a note to a lead.
func (s *LeadService) AddNote(ctx context.Context, leadID, authorID uuid.UUID, content string) error {
	// Verify lead exists.
	if _, err := s.repo.GetLeadByID(ctx, leadID); err != nil {
		return fmt.Errorf("lead not found: %w", err)
	}

	note := &domain.LeadNote{
		ID:        uuid.New(),
		LeadID:    leadID,
		AuthorID:  authorID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := s.repo.AddLeadNote(ctx, note); err != nil {
		return fmt.Errorf("failed to add note: %w", err)
	}

	log.Info().Str("lead_id", leadID.String()).Str("author_id", authorID.String()).Msg("note added to lead")
	return nil
}

// Assign reassigns a lead to a different agent and logs the event.
func (s *LeadService) Assign(ctx context.Context, leadID, agentID, actorID uuid.UUID) error {
	lead, err := s.repo.GetLeadByID(ctx, leadID)
	if err != nil {
		return fmt.Errorf("lead not found: %w", err)
	}

	previousAgentID := lead.AgentID
	lead.AgentID = &agentID
	lead.UpdatedAt = time.Now()

	if err := s.repo.UpdateLead(ctx, lead); err != nil {
		return fmt.Errorf("failed to assign lead: %w", err)
	}

	// Log the assignment event.
	if s.eventSvc != nil {
		payload := map[string]interface{}{
			"new_agent_id": agentID.String(),
		}
		if previousAgentID != nil {
			payload["previous_agent_id"] = previousAgentID.String()
		}
		_ = s.eventSvc.Log(ctx, actorID, "admin", "lead_assigned", "lead", leadID, payload)
	}

	log.Info().Str("lead_id", leadID.String()).Str("agent_id", agentID.String()).Msg("lead assigned to agent")
	return nil
}

// normalizePhone ensures the phone number has a +91 prefix if no country code is present.
func normalizePhone(phone string) string {
	phone = strings.TrimSpace(phone)
	if !strings.HasPrefix(phone, "+") {
		phone = "+91" + phone
	}
	return phone
}
