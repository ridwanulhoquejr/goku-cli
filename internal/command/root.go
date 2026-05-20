package command

import (
	"os"

	"github.com/ridwanulhoquejr/goku-cli/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "A CLI tool for config conversion and resource management",
	Long: `goku is a CLI toolkit built with Cobra.

Commands:
  goku convert  -i <file> -o <format> [-d <dir>]
  goku migrate  up|down
  goku save     -i <resource.json>
  goku list
  goku get      --id <id>
  goku update   --id <id> -i <resource.json>
  goku delete   --id <id>`,
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return db.Close()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
