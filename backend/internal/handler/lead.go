package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// LeadHandler handles lead HTTP endpoints.
type LeadHandler struct {
	svc *service.LeadService
}

// NewLeadHandler creates a new LeadHandler.
func NewLeadHandler(svc *service.LeadService) *LeadHandler {
	return &LeadHandler{svc: svc}
}

// Create handles POST /leads (public - investor submits interest).
func (h *LeadHandler) Create(c *fiber.Ctx) error {
	var input service.CreateLeadInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	if err := validator.ValidatePhone(input.Phone); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	lead, err := h.svc.Create(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "create_failed", err.Error())
	}

	return response.Created(c, lead)
}

// List handles GET /leads. Agents see their own leads, admins see all.
func (h *LeadHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	role := getUserRole(c)

	filters := service.LeadFilters{
		Status:    c.Query("status"),
		ProjectID: c.Query("project_id"),
		AgentID:   c.Query("agent_id"),
		Page:      queryInt(c, "page", 1),
		Limit:     queryInt(c, "limit", 20),
	}

	leads, total, err := h.svc.List(c.Context(), filters, role, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, leads, filters.Page, filters.Limit, total)
}

// GetByID handles GET /leads/:id.
func (h *LeadHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid lead ID")
	}

	lead, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "lead")
	}

	return response.Success(c, lead)
}

// UpdateStatus handles PUT /leads/:id/status.
func (h *LeadHandler) UpdateStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid lead ID")
	}

	var input service.UpdateLeadStatusInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	actorID := middleware.GetUserID(c)
	if err := h.svc.UpdateStatus(c.Context(), id, input.Status, input.Remarks, actorID); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "update_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "lead status updated"})
}

// AddNote handles POST /leads/:id/notes.
func (h *LeadHandler) AddNote(c *fiber.Ctx) error {
	idStr := c.Params("id")
	leadID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid lead ID")
	}

	var input service.AddNoteInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	authorID := middleware.GetUserID(c)
	if err := h.svc.AddNote(c.Context(), leadID, authorID, input.Content); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "note_failed", err.Error())
	}

	return response.Created(c, fiber.Map{"message": "note added"})
}

// Assign handles PUT /leads/:id/assign (admin only).
func (h *LeadHandler) Assign(c *fiber.Ctx) error {
	idStr := c.Params("id")
	leadID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid lead ID")
	}

	var input service.AssignLeadInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	agentID, err := uuid.Parse(input.AgentID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_agent_id", "invalid agent ID")
	}

	actorID := middleware.GetUserID(c)
	if err := h.svc.Assign(c.Context(), leadID, agentID, actorID); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "assign_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "lead assigned"})
}

// getUserRole extracts the authenticated user's role from Fiber locals.
func getUserRole(c *fiber.Ctx) string {
	role, _ := c.Locals(middleware.ContextKeyUserRole).(string)
	return role
}
