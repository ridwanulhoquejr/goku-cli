package command

import (
	"encoding/json"
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var saveInput string

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a resource file (JSON/YAML) into the database",
	Long: `Save reads a JSON or YAML resource file and inserts it into the database.

Examples:
  goku save -i server.json
  goku save -i database.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := db.Get()
		if err != nil {
			return err
		}

		resource, err := handlers.Save(conn, saveInput)
		if err != nil {
			return err
		}

		output, _ := json.MarshalIndent(resource, "", "  ")
		fmt.Printf("✔ Resource saved successfully:\n%s\n", string(output))
		return nil
	},
}

func init() {
	saveCmd.Flags().StringVarP(&saveInput, "input", "i", "", "Path to resource file (json or yaml)")
	saveCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(saveCmd)
}
