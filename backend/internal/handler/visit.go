package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// VisitHandler handles site visit HTTP endpoints.
type VisitHandler struct {
	svc *service.VisitService
}

// NewVisitHandler creates a new VisitHandler.
func NewVisitHandler(svc *service.VisitService) *VisitHandler {
	return &VisitHandler{svc: svc}
}

// Create handles POST /visits.
func (h *VisitHandler) Create(c *fiber.Ctx) error {
	agentID := middleware.GetUserID(c)
	if agentID == uuid.Nil {
		return response.Unauthorized(c)
	}

	var input service.CreateVisitInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	visit, err := h.svc.Create(c.Context(), agentID, input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "create_failed", err.Error())
	}

	return response.Created(c, visit)
}

// SubmitFeedback handles PUT /visits/:id/feedback.
func (h *VisitHandler) SubmitFeedback(c *fiber.Ctx) error {
	idStr := c.Params("id")
	visitID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid visit ID")
	}

	var input service.SubmitFeedbackInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if input.Rating != nil {
		if err := validator.ValidateRating(*input.Rating); err != nil {
			return response.ValidationError(c, validator.FormatValidationErrors(err))
		}
	}

	if err := h.svc.SubmitFeedback(c.Context(), visitID, input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "feedback_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "feedback submitted"})
}

// List handles GET /visits.
func (h *VisitHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	role := getUserRole(c)

	filters := service.VisitFilters{
		ProjectID: c.Query("project_id"),
		LeadID:    c.Query("lead_id"),
		Page:      queryInt(c, "page", 1),
		Limit:     queryInt(c, "limit", 20),
	}

	visits, total, err := h.svc.List(c.Context(), filters, role, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, visits, filters.Page, filters.Limit, total)
}
