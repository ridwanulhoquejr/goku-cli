package handlers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Delete removes a resource by its ID.
func Delete(conn *sqlx.DB, id int) error {
	query := `DELETE FROM resource_table WHERE id = $1`
	result, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	return nil
}
