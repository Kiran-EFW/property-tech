package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// AgentFilters holds the query filters for listing agents.
type AgentFilters struct {
	Tier     string `query:"tier"`
	IsActive *bool  `query:"is_active"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

// AgentPerformance holds performance metrics for an agent.
type AgentPerformance struct {
	AgentID        uuid.UUID `json:"agent_id"`
	TotalLeads     int       `json:"total_leads"`
	ActiveLeads    int       `json:"active_leads"`
	SiteVisits     int       `json:"site_visits"`
	Bookings       int       `json:"bookings"`
	ConversionRate float64   `json:"conversion_rate"` // bookings / total leads
	TotalEarnings  float64   `json:"total_earnings"`
	Period         string    `json:"period"`
}

// AgentRepository defines the database operations required by AgentService.
type AgentRepository interface {
	CreateAgent(ctx context.Context, agent *domain.Agent) error
	GetAgentByID(ctx context.Context, id uuid.UUID) (*domain.Agent, error)
	GetAgentByUserID(ctx context.Context, userID uuid.UUID) (*domain.Agent, error)
	ListAgents(ctx context.Context, filters AgentFilters) ([]*domain.Agent, int, error)
	UpdateAgent(ctx context.Context, agent *domain.Agent) error
	GetAgentPerformance(ctx context.Context, agentID uuid.UUID, period string) (*AgentPerformance, error)
}

// AgentService handles agent management business logic.
type AgentService struct {
	repo     AgentRepository
	eventSvc *EventService
}

// NewAgentService creates a new AgentService.
func NewAgentService(repo AgentRepository, eventSvc *EventService) *AgentService {
	return &AgentService{
		repo:     repo,
		eventSvc: eventSvc,
	}
}

// Register creates a new agent profile for an authenticated user.
func (s *AgentService) Register(ctx context.Context, userID uuid.UUID, reraNumber, pan, gst string) (*domain.Agent, error) {
	// Check if agent profile already exists.
	existing, err := s.repo.GetAgentByUserID(ctx, userID)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("agent profile already exists for this user")
	}

	agent := &domain.Agent{
		ID:        uuid.New(),
		UserID:    userID,
		Tier:      domain.AgentTierBronze,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if reraNumber != "" {
		agent.RERANumber = &reraNumber
	}
	if pan != "" {
		agent.PAN = &pan
	}
	if gst != "" {
		agent.GST = &gst
	}

	if err := s.repo.CreateAgent(ctx, agent); err != nil {
		return nil, fmt.Errorf("failed to create agent: %w", err)
	}

	log.Info().Str("agent_id", agent.ID.String()).Str("user_id", userID.String()).Msg("agent registered")
	return agent, nil
}

// List returns a filtered, paginated list of agents.
func (s *AgentService) List(ctx context.Context, filters AgentFilters) ([]*domain.Agent, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	agents, total, err := s.repo.ListAgents(ctx, filters)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list agents: %w", err)
	}

	return agents, total, nil
}

// GetPerformance returns performance metrics for an agent over a given period.
func (s *AgentService) GetPerformance(ctx context.Context, agentID uuid.UUID, period string) (*AgentPerformance, error) {
	if period == "" {
		period = "30d"
	}

	perf, err := s.repo.GetAgentPerformance(ctx, agentID, period)
	if err != nil {
		return nil, fmt.Errorf("failed to get performance: %w", err)
	}

	return perf, nil
}

// UpdateTier updates the tier of an agent and logs the event.
func (s *AgentService) UpdateTier(ctx context.Context, agentID uuid.UUID, tier string, actorID uuid.UUID) error {
	agent, err := s.repo.GetAgentByID(ctx, agentID)
	if err != nil {
		return fmt.Errorf("agent not found: %w", err)
	}

	previousTier := agent.Tier
	agent.Tier = domain.AgentTier(tier)
	agent.UpdatedAt = time.Now()

	if err := s.repo.UpdateAgent(ctx, agent); err != nil {
		return fmt.Errorf("failed to update agent tier: %w", err)
	}

	// Log the tier change event.
	if s.eventSvc != nil {
		_ = s.eventSvc.Log(ctx, actorID, "admin", "agent_tier_updated", "agent", agentID, map[string]interface{}{
			"previous_tier": string(previousTier),
			"new_tier":      tier,
		})
	}

	log.Info().Str("agent_id", agentID.String()).Str("tier", tier).Msg("agent tier updated")
	return nil
}
