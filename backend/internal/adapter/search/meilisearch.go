package search

import (
	"context"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog/log"
)

// SearchProvider defines the interface for full-text search operations.
type SearchProvider interface {
	SearchProviders(ctx context.Context, query string, filters SearchFilters) (*SearchResult, error)
	IndexProvider(ctx context.Context, doc ProviderDocument) error
}

// SearchFilters holds filter parameters for search queries.
type SearchFilters struct {
	CategoryID string
	Lat        float64
	Lng        float64
	RadiusKM   float64
}

// SearchResult holds the search response.
type SearchResult struct {
	Hits             []map[string]interface{} `json:"hits"`
	NbHits           int64                    `json:"nbHits"`
	ProcessingTimeMs int64                    `json:"processingTimeMs"`
}

// ProviderDocument represents a provider indexed in Meilisearch.
type ProviderDocument struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Skills     []string `json:"skills"`
	Postcode   string   `json:"postcode"`
	Location   string   `json:"location"`
	Rating     float64  `json:"rating"`
	Category   string   `json:"category"`
	CategoryID string   `json:"category_id"`
	Language   []string `json:"language"`
	Geo        *GeoPoint `json:"_geo,omitempty"`
}

// GeoPoint represents a geographic coordinate for Meilisearch geo search.
type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// MeiliSearchProvider implements SearchProvider using Meilisearch.
type MeiliSearchProvider struct {
	client meilisearch.ServiceManager
}

// NewMeiliSearchProvider creates a new Meilisearch-backed search provider.
func NewMeiliSearchProvider(client meilisearch.ServiceManager) *MeiliSearchProvider {
	return &MeiliSearchProvider{client: client}
}

// SearchProviders searches for providers matching the query and filters.
func (m *MeiliSearchProvider) SearchProviders(ctx context.Context, query string, filters SearchFilters) (*SearchResult, error) {
	_ = ctx // Meilisearch Go SDK does not support context yet

	searchReq := &meilisearch.SearchRequest{
		Limit: 20,
	}

	// Build filter string.
	var filterParts []string
	if filters.CategoryID != "" {
		filterParts = append(filterParts, fmt.Sprintf("category_id = %q", filters.CategoryID))
	}

	if len(filterParts) > 0 {
		combined := ""
		for i, part := range filterParts {
			if i > 0 {
				combined += " AND "
			}
			combined += part
		}
		filter := []string{combined}
		searchReq.Filter = filter
	}

	// Add geo search if coordinates provided.
	if filters.Lat != 0 && filters.Lng != 0 && filters.RadiusKM > 0 {
		radiusMeters := int(filters.RadiusKM * 1000)
		searchReq.Filter = append(searchReq.Filter.([]string),
			fmt.Sprintf("_geoRadius(%f, %f, %d)", filters.Lat, filters.Lng, radiusMeters))
	}

	resp, err := m.client.Index("providers").Search(query, searchReq)
	if err != nil {
		return nil, fmt.Errorf("meilisearch search: %w", err)
	}

	hits := make([]map[string]interface{}, len(resp.Hits))
	for i, hit := range resp.Hits {
		if m, ok := hit.(map[string]interface{}); ok {
			hits[i] = m
		}
	}

	log.Debug().
		Str("query", query).
		Int("hits", len(hits)).
		Int64("processing_ms", resp.ProcessingTimeMs).
		Msg("meilisearch search completed")

	return &SearchResult{
		Hits:             hits,
		NbHits:           resp.EstimatedTotalHits,
		ProcessingTimeMs: resp.ProcessingTimeMs,
	}, nil
}

// IndexProvider adds or updates a provider document in the search index.
func (m *MeiliSearchProvider) IndexProvider(ctx context.Context, doc ProviderDocument) error {
	_ = ctx

	_, err := m.client.Index("providers").AddDocuments([]ProviderDocument{doc}, "id")
	if err != nil {
		return fmt.Errorf("meilisearch index provider: %w", err)
	}

	log.Debug().
		Str("provider_id", doc.ID).
		Str("name", doc.Name).
		Msg("provider indexed in meilisearch")

	return nil
}
