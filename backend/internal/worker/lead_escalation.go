package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// LeadEscalationPayload is the JSON payload for lead escalation tasks.
type LeadEscalationPayload struct {
	LeadID  string `json:"lead_id"`
	AgentID string `json:"agent_id"`
}

// LeadEscalationRepository defines the data access needed for lead escalation.
type LeadEscalationRepository interface {
	GetUncontactedLeadsOlderThan(ctx context.Context, threshold time.Duration) ([]*domain.Lead, error)
	GetAgentWithFewestActiveLeads(ctx context.Context) (*domain.Agent, error)
	UpdateLead(ctx context.Context, lead *domain.Lead) error
}

// LeadEscalationHandler processes lead escalation tasks.
type LeadEscalationHandler struct {
	repo LeadEscalationRepository
}

// NewLeadEscalationHandler creates a new LeadEscalationHandler.
func NewLeadEscalationHandler(repo LeadEscalationRepository) *LeadEscalationHandler {
	return &LeadEscalationHandler{repo: repo}
}

// HandleLeadEscalation checks for uncontacted leads older than 5 minutes
// and reassigns them to the next available agent.
func (h *LeadEscalationHandler) HandleLeadEscalation(ctx context.Context, task *asynq.Task) error {
	log.Info().Str("task_type", task.Type()).Msg("worker: processing lead escalation")

	// Find leads that are still in "new" status and older than 5 minutes.
	threshold := 5 * time.Minute
	leads, err := h.repo.GetUncontactedLeadsOlderThan(ctx, threshold)
	if err != nil {
		return fmt.Errorf("lead_escalation: failed to get uncontacted leads: %w", err)
	}

	if len(leads) == 0 {
		log.Info().Msg("lead_escalation: no uncontacted leads to escalate")
		return nil
	}

	log.Info().Int("count", len(leads)).Msg("lead_escalation: found uncontacted leads to reassign")

	for _, lead := range leads {
		// Find the next available agent (with fewest active leads).
		agent, err := h.repo.GetAgentWithFewestActiveLeads(ctx)
		if err != nil {
			log.Error().Err(err).Str("lead_id", lead.ID.String()).Msg("lead_escalation: failed to find available agent")
			continue
		}
		if agent == nil {
			log.Warn().Str("lead_id", lead.ID.String()).Msg("lead_escalation: no available agents")
			continue
		}

		// Skip if the lead is already assigned to this agent.
		if lead.AgentID != nil && *lead.AgentID == agent.ID {
			continue
		}

		previousAgentID := lead.AgentID
		lead.AgentID = &agent.ID
		lead.UpdatedAt = time.Now()

		if err := h.repo.UpdateLead(ctx, lead); err != nil {
			log.Error().Err(err).
				Str("lead_id", lead.ID.String()).
				Str("new_agent_id", agent.ID.String()).
				Msg("lead_escalation: failed to reassign lead")
			continue
		}

		prevStr := "none"
		if previousAgentID != nil {
			prevStr = previousAgentID.String()
		}

		log.Info().
			Str("lead_id", lead.ID.String()).
			Str("previous_agent", prevStr).
			Str("new_agent", agent.ID.String()).
			Msg("lead_escalation: lead reassigned")
	}

	return nil
}

// NewLeadEscalationTask creates a new Asynq task for lead escalation.
func NewLeadEscalationTask() (*asynq.Task, error) {
	return asynq.NewTask(TaskLeadEscalation, nil), nil
}

// NewLeadEscalationTaskForLead creates a task for a specific lead escalation.
func NewLeadEscalationTaskForLead(leadID, agentID uuid.UUID) (*asynq.Task, error) {
	payload := LeadEscalationPayload{
		LeadID:  leadID.String(),
		AgentID: agentID.String(),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal lead escalation payload: %w", err)
	}
	return asynq.NewTask(TaskLeadEscalation, data), nil
}
