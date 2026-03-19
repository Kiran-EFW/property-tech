package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// AgentHandler handles agent HTTP endpoints.
type AgentHandler struct {
	svc *service.AgentService
}

// NewAgentHandler creates a new AgentHandler.
func NewAgentHandler(svc *service.AgentService) *AgentHandler {
	return &AgentHandler{svc: svc}
}

// registerAgentRequest is the request body for POST /agents/register.
type registerAgentRequest struct {
	RERANumber string `json:"rera_number"`
	PAN        string `json:"pan"`
	GST        string `json:"gst"`
}

// Register handles POST /agents/register.
func (h *AgentHandler) Register(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	var req registerAgentRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	agent, err := h.svc.Register(c.Context(), userID, req.RERANumber, req.PAN, req.GST)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "registration_failed", err.Error())
	}

	return response.Created(c, agent)
}

// List handles GET /agents (admin only).
func (h *AgentHandler) List(c *fiber.Ctx) error {
	filters := service.AgentFilters{
		Tier: c.Query("tier"),
		Page: queryInt(c, "page", 1),
		Limit: queryInt(c, "limit", 20),
	}

	agents, total, err := h.svc.List(c.Context(), filters)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, agents, filters.Page, filters.Limit, total)
}

// GetPerformance handles GET /agents/:id/performance.
func (h *AgentHandler) GetPerformance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	agentID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid agent ID")
	}

	period := c.Query("period", "30d")

	perf, err := h.svc.GetPerformance(c.Context(), agentID, period)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "performance_failed", err.Error())
	}

	return response.Success(c, perf)
}

// updateTierRequest is the request body for PUT /agents/:id/tier.
type updateTierRequest struct {
	Tier string `json:"tier" validate:"required,oneof=bronze silver gold platinum"`
}

// UpdateTier handles PUT /agents/:id/tier (admin only).
func (h *AgentHandler) UpdateTier(c *fiber.Ctx) error {
	idStr := c.Params("id")
	agentID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid agent ID")
	}

	var req updateTierRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(req); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	actorID := middleware.GetUserID(c)
	if err := h.svc.UpdateTier(c.Context(), agentID, req.Tier, actorID); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "update_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "agent tier updated"})
}
