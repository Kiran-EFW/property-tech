package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// ProjectHandler handles project HTTP endpoints.
type ProjectHandler struct {
	svc *service.ProjectService
}

// NewProjectHandler creates a new ProjectHandler.
func NewProjectHandler(svc *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

// List handles GET /projects.
func (h *ProjectHandler) List(c *fiber.Ctx) error {
	filters := service.ProjectFilters{
		Status: c.Query("status"),
		City:   c.Query("city"),
		Page:   queryInt(c, "page", 1),
		Limit:  queryInt(c, "limit", 20),
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		if v, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filters.MinPrice = v
		}
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if v, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filters.MaxPrice = v
		}
	}

	projects, total, err := h.svc.List(c.Context(), filters)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, projects, filters.Page, filters.Limit, total)
}

// GetBySlug handles GET /projects/:slug.
func (h *ProjectHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return response.Error(c, fiber.StatusBadRequest, "missing_slug", "project slug is required")
	}

	project, err := h.svc.GetBySlug(c.Context(), slug)
	if err != nil {
		return response.NotFound(c, "project")
	}

	return response.Success(c, project)
}

// Create handles POST /projects (admin only).
func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	var input service.CreateProjectInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(input); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	project, err := h.svc.Create(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "create_failed", err.Error())
	}

	return response.Created(c, project)
}

// Update handles PUT /projects/:id (admin only).
func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid project ID")
	}

	var input service.UpdateProjectInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	project, err := h.svc.Update(c.Context(), id, input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "update_failed", err.Error())
	}

	return response.Success(c, project)
}

// GetInventory handles GET /projects/:id/inventory.
func (h *ProjectHandler) GetInventory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid project ID")
	}

	units, err := h.svc.GetInventory(c.Context(), id)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "inventory_failed", err.Error())
	}

	return response.Success(c, units)
}

// GetDueDiligence handles GET /projects/:id/due-diligence.
func (h *ProjectHandler) GetDueDiligence(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_id", "invalid project ID")
	}

	report, err := h.svc.GetDueDiligence(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "due-diligence report")
	}

	return response.Success(c, report)
}

// queryInt extracts an integer query parameter with a default value.
func queryInt(c *fiber.Ctx, key string, defaultVal int) int {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return v
}
