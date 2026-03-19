package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meilisearch/meilisearch-go"

	"github.com/proptech/backend/pkg/response"
)

// SearchHandler handles search HTTP endpoints backed by Meilisearch.
type SearchHandler struct {
	client meilisearch.ServiceManager
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler(client meilisearch.ServiceManager) *SearchHandler {
	return &SearchHandler{client: client}
}

// SearchProjects handles GET /search/projects.
func (h *SearchHandler) SearchProjects(c *fiber.Ctx) error {
	query := c.Query("q", "")
	page := queryInt(c, "page", 1)
	limit := queryInt(c, "limit", 20)

	if h.client == nil {
		return response.Error(c, fiber.StatusServiceUnavailable, "search_unavailable", "search service is not configured")
	}

	offset := (page - 1) * limit
	searchReq := &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	}

	resp, err := h.client.Index("projects").Search(query, searchReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "search_failed", fmt.Sprintf("search failed: %v", err))
	}

	hits := make([]map[string]interface{}, 0, len(resp.Hits))
	for _, hit := range resp.Hits {
		if m, ok := hit.(map[string]interface{}); ok {
			hits = append(hits, m)
		}
	}

	return response.Paginated(c, hits, page, limit, int(resp.EstimatedTotalHits))
}

// SearchAreas handles GET /search/areas.
func (h *SearchHandler) SearchAreas(c *fiber.Ctx) error {
	query := c.Query("q", "")
	page := queryInt(c, "page", 1)
	limit := queryInt(c, "limit", 20)

	if h.client == nil {
		return response.Error(c, fiber.StatusServiceUnavailable, "search_unavailable", "search service is not configured")
	}

	offset := (page - 1) * limit
	searchReq := &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	}

	resp, err := h.client.Index("areas").Search(query, searchReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "search_failed", fmt.Sprintf("search failed: %v", err))
	}

	hits := make([]map[string]interface{}, 0, len(resp.Hits))
	for _, hit := range resp.Hits {
		if m, ok := hit.(map[string]interface{}); ok {
			hits = append(hits, m)
		}
	}

	return response.Paginated(c, hits, page, limit, int(resp.EstimatedTotalHits))
}
