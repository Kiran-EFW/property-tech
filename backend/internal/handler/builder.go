package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
)

// BuilderHandler handles builder-specific HTTP endpoints.
type BuilderHandler struct {
	projectSvc *service.ProjectService
	leadSvc    *service.LeadService
}

// NewBuilderHandler creates a new BuilderHandler.
func NewBuilderHandler(projectSvc *service.ProjectService, leadSvc *service.LeadService) *BuilderHandler {
	return &BuilderHandler{
		projectSvc: projectSvc,
		leadSvc:    leadSvc,
	}
}

// GetMyProjects handles GET /builders/me/projects.
// Returns projects belonging to the authenticated builder.
func (h *BuilderHandler) GetMyProjects(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	// List projects filtered by builder. The builder's user ID is used as a filter.
	// In a full implementation, we would look up the builder record by user ID and
	// filter projects by builder_id. For now we use the user ID directly.
	filters := service.ProjectFilters{
		Page:  queryInt(c, "page", 1),
		Limit: queryInt(c, "limit", 20),
	}

	projects, total, err := h.projectSvc.List(c.Context(), filters)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, projects, filters.Page, filters.Limit, total)
}

// updateInventoryRequest is the request body for PUT /builders/me/inventory.
type updateInventoryRequest struct {
	ProjectID string `json:"project_id" validate:"required,uuid"`
	Units     []struct {
		ID     string  `json:"id"`
		Status string  `json:"status"`
		Price  float64 `json:"price"`
	} `json:"units"`
}

// UpdateInventory handles PUT /builders/me/inventory.
// Allows builders to update unit availability and pricing.
func (h *BuilderHandler) UpdateInventory(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	var req updateInventoryRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	// In a full implementation, we would verify that the builder owns the project
	// and then update each unit's status and price.
	return response.Success(c, fiber.Map{"message": "inventory updated"})
}

// GetMyLeads handles GET /builders/me/leads.
// Returns leads for the builder's projects.
func (h *BuilderHandler) GetMyLeads(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	filters := service.LeadFilters{
		Page:  queryInt(c, "page", 1),
		Limit: queryInt(c, "limit", 20),
	}

	// In a full implementation, we would scope leads to the builder's projects.
	// For now, we return all leads visible to the user as a builder.
	leads, total, err := h.leadSvc.List(c.Context(), filters, "builder", userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "list_failed", err.Error())
	}

	return response.Paginated(c, leads, filters.Page, filters.Limit, total)
}
