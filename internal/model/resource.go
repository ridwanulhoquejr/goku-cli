package model

import "time"

type Resource struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	RawData   []byte    `json:"-" db:"data"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Data holds the unmarshalled JSONB content for display.
	Data map[string]interface{} `json:"data" db:"-"`
}
