package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ridwanulhoquejr/goku-cli/internal/model"
)

// Get retrieves a single resource by its ID.
func Get(conn *sqlx.DB, id int) (*model.Resource, error) {
	query := `
		SELECT id, name, type, data, created_at, updated_at
		FROM resource_table
		WHERE id = $1
	`
	var resource model.Resource
	if err := conn.QueryRowx(query, id).StructScan(&resource); err != nil {
		return nil, fmt.Errorf("resource with id %d not found: %w", id, err)
	}

	if err := json.Unmarshal(resource.RawData, &resource.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &resource, nil
}
