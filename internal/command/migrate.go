package command

import (
	"fmt"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/spf13/cobra"
)

var migrationsPath string

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := db.RunMigrations(migrationsPath); err != nil {
			return err
		}
		fmt.Println("✔ Migrations applied successfully.")
		return nil
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := db.RollbackMigrations(migrationsPath); err != nil {
			return err
		}
		fmt.Println("✔ Migration rolled back successfully.")
		return nil
	},
}

func init() {
	migrateCmd.PersistentFlags().StringVar(&migrationsPath, "path", "file://migrations", "Path to migration files")

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateCmd)
}
