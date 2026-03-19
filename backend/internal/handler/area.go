package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// AreaHandler handles area/micro-market HTTP endpoints.
type AreaHandler struct {
	svc *service.AreaService
}

// NewAreaHandler creates a new AreaHandler.
func NewAreaHandler(svc *service.AreaService) *AreaHandler {
	return &AreaHandler{svc: svc}
}

// List handles GET /areas.
func (h *AreaHandler) List(c *fiber.Ctx) error {
	filters := service.AreaFilters{
		City:  c.Query("city"),
		State: c.Query("state"),
		Page:  queryInt(c, "page", 1),
		Limit: queryInt(c, "limit", 20),
	}

	areas, total, err := h.svc.List(c.Context(), filters)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, areas, filters.Page, filters.Limit, total)
}

// GetBySlug handles GET /areas/:slug.
func (h *AreaHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return response.Error(c, fiber.StatusBadRequest, "missing_slug", "area slug is required")
	}

	area, err := h.svc.GetBySlug(c.Context(), slug)
	if err != nil {
		return response.NotFound(c, "area")
	}

	return response.Success(c, area)
}

// Create handles POST /areas (admin only).
func (h *AreaHandler) Create(c *fiber.Ctx) error {
	var input service.CreateAreaInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	area, err := h.svc.Create(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "create_failed", err.Error())
	}

	return response.Created(c, area)
}
