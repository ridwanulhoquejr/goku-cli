package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ridwanulhoquejr/goku-cli/internal/model"
)

// Save reads a JSON/YAML file and inserts it into the database.
func Save(conn *sqlx.DB, filePath string) (*model.Resource, error) {
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
		INSERT INTO resource_table (name, type, data)
		VALUES ($1, $2, $3)
		RETURNING id, name, type, data, created_at, updated_at
	`
	var resource model.Resource
	if err := conn.QueryRowx(query, name, format, jsonData).StructScan(&resource); err != nil {
		return nil, fmt.Errorf("failed to insert resource: %w", err)
	}

	if err := json.Unmarshal(resource.RawData, &resource.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal returned data: %w", err)
	}

	return &resource, nil
}
