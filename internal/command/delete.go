package command

import (
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := db.Get()
		if err != nil {
			return err
		}

		if err := handlers.Delete(conn, deleteID); err != nil {
			return err
		}

		fmt.Printf("✔ Resource with id %d deleted successfully.\n", deleteID)
		return nil
	},
}

func init() {
	deleteCmd.Flags().IntVar(&deleteID, "id", 0, "Resource ID to delete")
	deleteCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(deleteCmd)
}
