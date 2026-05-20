package command

import (
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all resources from the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := db.Get()
		if err != nil {
			return err
		}

		resources, err := handlers.List(conn)
		if err != nil {
			return err
		}

		if len(resources) == 0 {
			fmt.Println("No resources found.")
			return nil
		}

		fmt.Printf("%-5s %-20s %-8s %-25s\n", "ID", "NAME", "TYPE", "CREATED AT")
		fmt.Printf("%-5s %-20s %-8s %-25s\n", "---", "----", "----", "----------")
		for _, r := range resources {
			fmt.Printf("%-5d %-20s %-8s %-25s\n",
				r.ID, r.Name, r.Type, r.CreatedAt.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf("\nTotal: %d resource(s)\n", len(resources))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
