package search

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

// MeiliClient wraps the Meilisearch SDK and provides index initialization.
type MeiliClient struct {
	client meilisearch.ServiceManager
}

// NewMeiliClient creates a new MeiliClient and verifies the connection is healthy.
func NewMeiliClient(url, apiKey string) (*MeiliClient, error) {
	client := meilisearch.New(url, meilisearch.WithAPIKey(apiKey))
	// Verify connection
	if _, err := client.Health(); err != nil {
		return nil, fmt.Errorf("meilisearch health check failed: %w", err)
	}
	return &MeiliClient{client: client}, nil
}

// InitIndexes creates or updates the required search indexes with proper settings.
func (m *MeiliClient) InitIndexes() error {
	indexes := []struct {
		uid        string
		primaryKey string
		searchable []string
		filterable []string
		sortable   []string
	}{
		{
			uid:        "projects",
			primaryKey: "id",
			searchable: []string{"name", "description", "builder_name", "area_name", "slug"},
			filterable: []string{"builder_id", "status", "area_id", "price_min", "price_max", "carpet_area_min", "carpet_area_max"},
			sortable:   []string{"price_min", "created_at", "carpet_area_min"},
		},
		{
			uid:        "builders",
			primaryKey: "id",
			searchable: []string{"name", "rera_number"},
			filterable: []string{"is_active"},
			sortable:   []string{"name", "created_at"},
		},
		{
			uid:        "areas",
			primaryKey: "id",
			searchable: []string{"name", "slug", "description"},
			filterable: []string{},
			sortable:   []string{"name"},
		},
	}

	for _, idx := range indexes {
		// Create index if it doesn't exist
		taskInfo, err := m.client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        idx.uid,
			PrimaryKey: idx.primaryKey,
		})
		if err != nil {
			// Index may already exist, that's fine
			_ = taskInfo
		}

		index := m.client.Index(idx.uid)

		// Update searchable attributes
		if _, err := index.UpdateSearchableAttributes(&idx.searchable); err != nil {
			return fmt.Errorf("update searchable attrs for %s: %w", idx.uid, err)
		}

		// Update filterable attributes (SDK requires []interface{})
		filterableIface := make([]interface{}, len(idx.filterable))
		for i, v := range idx.filterable {
			filterableIface[i] = v
		}
		if _, err := index.UpdateFilterableAttributes(&filterableIface); err != nil {
			return fmt.Errorf("update filterable attrs for %s: %w", idx.uid, err)
		}

		// Update sortable attributes
		if _, err := index.UpdateSortableAttributes(&idx.sortable); err != nil {
			return fmt.Errorf("update sortable attrs for %s: %w", idx.uid, err)
		}
	}

	return nil
}

// Client returns the underlying meilisearch client for use by services.
func (m *MeiliClient) Client() meilisearch.ServiceManager {
	return m.client
}
