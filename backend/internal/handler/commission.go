package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
)

// CommissionHandler handles commission HTTP endpoints.
type CommissionHandler struct {
	svc *service.CommissionService
}

// NewCommissionHandler creates a new CommissionHandler.
func NewCommissionHandler(svc *service.CommissionService) *CommissionHandler {
	return &CommissionHandler{svc: svc}
}

// List handles GET /commissions.
func (h *CommissionHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	role := getUserRole(c)

	filters := service.CommissionFilters{
		Status:  c.Query("status"),
		AgentID: c.Query("agent_id"),
		Page:    queryInt(c, "page", 1),
		Limit:   queryInt(c, "limit", 20),
	}

	commissions, total, err := h.svc.List(c.Context(), filters, role, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, commissions, filters.Page, filters.Limit, total)
}

// Approve handles POST /commissions/:id/approve (admin only).
func (h *CommissionHandler) Approve(c *fiber.Ctx) error {
	idStr := c.Params("id")
	commissionID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid commission ID")
	}

	actorID := middleware.GetUserID(c)
	if err := h.svc.Approve(c.Context(), commissionID, actorID); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "approve_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "commission approved"})
}

// GetSummary handles GET /commissions/summary.
func (h *CommissionHandler) GetSummary(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	role := getUserRole(c)

	summary, err := h.svc.GetSummary(c.Context(), role, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "summary_failed", err.Error())
	}

	return response.Success(c, summary)
}
