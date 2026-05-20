package command

import (
	"encoding/json"
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var (
	updateID    int
	updateInput string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a resource by ID with new data from a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := db.Get()
		if err != nil {
			return err
		}

		resource, err := handlers.Update(conn, updateID, updateInput)
		if err != nil {
			return err
		}

		output, _ := json.MarshalIndent(resource, "", "  ")
		fmt.Printf("✔ Resource updated successfully:\n%s\n", string(output))
		return nil
	},
}

func init() {
	updateCmd.Flags().IntVar(&updateID, "id", 0, "Resource ID to update")
	updateCmd.Flags().StringVarP(&updateInput, "input", "i", "", "Path to the new resource file")

	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(updateCmd)
}
