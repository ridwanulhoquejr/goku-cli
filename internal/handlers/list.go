package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ridwanulhoquejr/goku-cli/internal/model"
)

// List retrieves all resources from the database.
func List(conn *sqlx.DB) ([]model.Resource, error) {
	query := `
		SELECT id, name, type, data, created_at, updated_at
		FROM resource_table
		ORDER BY id ASC
	`
	var resources []model.Resource
	if err := conn.Select(&resources, query); err != nil {
		return nil, fmt.Errorf("failed to list resources: %w", err)
	}

	for i := range resources {
		if err := json.Unmarshal(resources[i].RawData, &resources[i].Data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data for resource %d: %w", resources[i].ID, err)
		}
	}

	return resources, nil
}
