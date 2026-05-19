package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ridwanulhoquejr/goku-cli/internal/model"
)

// Update reads a new file and replaces the data of an existing resource.
func Update(conn *sqlx.DB, id int, filePath string) (*model.Resource, error) {
	format, err := DetectFormat(filePath)
	if err != nil {
		return nil, err
	}

	data, err := ReadFile(filePath, format)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	name := DeriveResourceName(filePath)

	query := `
		UPDATE resource_table
		SET name = $1, type = $2, data = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, type, data, created_at, updated_at
	`
	var resource model.Resource
	if err := conn.QueryRowx(query, name, format, jsonData, id).StructScan(&resource); err != nil {
		return nil, fmt.Errorf("failed to update resource with id %d: %w", id, err)
	}

	if err := json.Unmarshal(resource.RawData, &resource.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal returned data: %w", err)
	}

	return &resource, nil
}
