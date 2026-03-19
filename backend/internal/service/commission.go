package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/domain"
)

// CommissionFilters holds the query filters for listing commissions.
type CommissionFilters struct {
	Status  string `query:"status"`
	AgentID string `query:"agent_id"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}

// CommissionSummary holds aggregated commission statistics.
type CommissionSummary struct {
	TotalEarned   float64 `json:"total_earned"`
	TotalPending  float64 `json:"total_pending"`
	TotalApproved float64 `json:"total_approved"`
	TotalPaid     float64 `json:"total_paid"`
	Count         int     `json:"count"`
}

// TierBonusMultiplier returns the bonus multiplier for a given agent tier.
var TierBonusMultiplier = map[domain.AgentTier]float64{
	domain.AgentTierBronze:   1.0,
	domain.AgentTierSilver:   1.1,
	domain.AgentTierGold:     1.25,
	domain.AgentTierPlatinum: 1.5,
}

// CommissionRepository defines the database operations required by CommissionService.
type CommissionRepository interface {
	CreateCommission(ctx context.Context, commission *domain.Commission) error
	GetCommissionByID(ctx context.Context, id uuid.UUID) (*domain.Commission, error)
	ListCommissions(ctx context.Context, filters CommissionFilters) ([]*domain.Commission, int, error)
	ListCommissionsByAgent(ctx context.Context, agentID uuid.UUID, filters CommissionFilters) ([]*domain.Commission, int, error)
	UpdateCommission(ctx context.Context, commission *domain.Commission) error
	GetCommissionSummary(ctx context.Context, agentID *uuid.UUID) (*CommissionSummary, error)
}

// CommissionService handles commission business logic.
type CommissionService struct {
	repo     CommissionRepository
	eventSvc *EventService
}

// NewCommissionService creates a new CommissionService.
func NewCommissionService(repo CommissionRepository, eventSvc *EventService) *CommissionService {
	return &CommissionService{
		repo:     repo,
		eventSvc: eventSvc,
	}
}

// CalculateCommission computes the commission amount for a booking.
// Formula: base rate from builder * agent tier bonus.
// TDS is deducted at 5% of the gross amount.
func CalculateCommission(agreementValue, baseRatePercent float64, tier domain.AgentTier) (grossAmount, tds, netAmount float64) {
	multiplier, ok := TierBonusMultiplier[tier]
	if !ok {
		multiplier = 1.0
	}

	grossAmount = agreementValue * (baseRatePercent / 100.0) * multiplier
	tds = grossAmount * 0.05 // 5% TDS
	netAmount = grossAmount - tds
	return
}

// List returns commissions scoped by role.
func (s *CommissionService) List(ctx context.Context, filters CommissionFilters, role string, userID uuid.UUID) ([]*domain.Commission, int, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Limit < 1 || filters.Limit > 100 {
		filters.Limit = 20
	}

	if role == "admin" || role == "super_admin" {
		return s.repo.ListCommissions(ctx, filters)
	}

	return s.repo.ListCommissionsByAgent(ctx, userID, filters)
}

// Approve transitions a commission to the approved state.
func (s *CommissionService) Approve(ctx context.Context, commissionID, actorID uuid.UUID) error {
	commission, err := s.repo.GetCommissionByID(ctx, commissionID)
	if err != nil {
		return fmt.Errorf("commission not found: %w", err)
	}

	if commission.Status != domain.CommissionStatusPending {
		return fmt.Errorf("commission is not in pending status")
	}

	commission.Status = domain.CommissionStatusApproved
	commission.UpdatedAt = time.Now()

	if err := s.repo.UpdateCommission(ctx, commission); err != nil {
		return fmt.Errorf("failed to approve commission: %w", err)
	}

	// Log the approval event.
	if s.eventSvc != nil {
		_ = s.eventSvc.Log(ctx, actorID, "admin", "commission_approved", "commission", commissionID, map[string]interface{}{
			"amount":     commission.Amount,
			"net_amount": commission.NetAmount,
		})
	}

	log.Info().Str("commission_id", commissionID.String()).Msg("commission approved")
	return nil
}

// GetSummary returns aggregated commission statistics.
// If role is not admin/super_admin, scopes to the user's agent commissions.
func (s *CommissionService) GetSummary(ctx context.Context, role string, userID uuid.UUID) (*CommissionSummary, error) {
	var agentID *uuid.UUID
	if role != "admin" && role != "super_admin" {
		agentID = &userID
	}

	summary, err := s.repo.GetCommissionSummary(ctx, agentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get commission summary: %w", err)
	}

	return summary, nil
}
