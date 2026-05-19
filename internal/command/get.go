package command

import (
	"encoding/json"
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var getID int

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := db.Get()
		if err != nil {
			return err
		}

		resource, err := handlers.Get(conn, getID)
		if err != nil {
			return err
		}

		output, _ := json.MarshalIndent(resource, "", "  ")
		fmt.Printf("✔ Resource found:\n%s\n", string(output))
		return nil
	},
}

func init() {
	getCmd.Flags().IntVar(&getID, "id", 0, "Resource ID to retrieve")
	getCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(getCmd)
}
