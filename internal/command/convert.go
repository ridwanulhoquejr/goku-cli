package command

import (
	"fmt"
	"strings"

	"github.com/ridwanulhoquejr/goku-cli/internal/handlers"
	"github.com/spf13/cobra"
)

var (
	convertInput  string
	convertOutput string
	convertDir    string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert between JSON and YAML configuration files",
	Long: `Convert reads a JSON or YAML file and outputs it in the other format.

Examples:
  goku convert -i config.json -o yaml
  goku convert -i config.yaml -o json -d ./output`,
	RunE: func(cmd *cobra.Command, args []string) error {
		outFmt := strings.ToLower(convertOutput)
		if outFmt != "json" && outFmt != "yaml" {
			return fmt.Errorf("invalid output format %q (must be json or yaml)", convertOutput)
		}

		outputPath, err := handlers.Convert(convertInput, outFmt, convertDir)
		if err != nil {
			return err
		}

		fmt.Printf("✔ Converted successfully: %s\n", outputPath)
		return nil
	},
}

func init() {
	convertCmd.Flags().StringVarP(&convertInput, "input", "i", "", "Path to input file (json or yaml)")
	convertCmd.Flags().StringVarP(&convertOutput, "output", "o", "", "Output format: json or yaml")
	convertCmd.Flags().StringVarP(&convertDir, "dir", "d", "", "Output directory (defaults to system temp dir)")

	convertCmd.MarkFlagRequired("input")
	convertCmd.MarkFlagRequired("output")

	rootCmd.AddCommand(convertCmd)
}
