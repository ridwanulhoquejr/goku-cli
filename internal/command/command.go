package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI flags variables
var (
	inputFile    string
	outputFormat string
	outputFile   string
)

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "A CLI tool to convert between JSON to YAML configuration files",
	Long: `goku is a configuration file converter built with Cobra.
It reads a JSON or YAML file, maps it into a Go data structure,
and outputs it in the requested format.

Usage:
  goku -i <file_path> -o <json|yaml> [-d <-output_dir>]

Examples:
  goku -i config.json -o yaml
  goku -i config.yaml -o json -d ./output`,
	RunE: runConvert,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error in root-command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to input file (json or yaml)")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format: json or yaml")
	rootCmd.Flags().StringVarP(&outputFile, "dir", "d", "", "Output directory (defaults to system temp dir)")
}
