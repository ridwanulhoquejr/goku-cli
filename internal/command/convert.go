package command

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// it will take a filepath as string and recognize the file extension
// wheather its a .json or .yaml file
func extractFormat(filePath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".json":
		return "json", nil
	case ".yaml", ".yml":
		return "yaml", nil
	default:
		return "", fmt.Errorf("Unsupported file extension %q (supported: .json, .yaml)", ext)
	}
}

// it will take a filepath and extracted format
// converted the given fileformat to Go datastructure (map)
func readFile(filePath string, format string) (map[string]interface{}, error) {
	// first, we need to read the input file via os
	// which will return []bytes and an error
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// we will store the Unmarshalled json object to this variable
	var result map[string]interface{}

	// now check by switch cases what ext is it then proceed with that unmarshal
	switch format {
	case "json":
		if err := json.Unmarshal(data, &result); err != nil {
			return result, fmt.Errorf("failed to unmarshal json file: %w", err)
		}
	case "yaml":
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("failed to unamrshal yaml file: %w", err)
		}
		// here we dont need any default case since we already checks while extracting file type beforehaand
	}
	return result, nil
}

// buildOutputPath constructs the output file path.
// Uses the provided dir, or falls back to the system temp directory.
// The filename is derived from the input file with the new extension.
func buildOutputPath(inputPath string, outFmt string, dir string) (string, error) {
	// If no directory provided, use system temp dir
	if dir == "" {
		dir = os.TempDir()
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory %q: %w", dir, err)
	}

	// Take input filename and swap the extension
	// e.g. config.json → config.yaml
	baseName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	newExt := "." + outFmt
	outputFile := filepath.Join(dir, baseName+newExt)

	return outputFile, nil
}

// writeOutput marshals the data into the target format and writes to a file.
func writeOutput(data map[string]interface{}, format string, outputPath string) error {

	var output []byte
	var err error

	switch format {
	case "json":
		output, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to convert to JSON: %w", err)
		}
		output = append(output, '\n') // trailing newline for JSON
	case "yaml":
		output, err = yaml.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to convert to YAML: %w", err)
		}
	}

	// Write to file with 0644 permissions (owner read/write, others read)
	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Converted successfully: %s\n", outputPath)
	return nil
}

// runConvert is the main command handler.
func runConvert(cmd *cobra.Command, args []string) error {
	// Validate output format
	outFmt := strings.ToLower(outputFormat)
	if outFmt != "json" && outFmt != "yaml" {
		return fmt.Errorf("invalid output format %q (must be json or yaml)", outputFormat)
	}

	// Check if file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file %q does not exist", inputFile)
	}

	// Detect input format from extension
	inFmt, err := extractFormat(inputFile)
	if err != nil {
		return err
	}

	// Edge case: input and output formats are the same
	if inFmt == outFmt {
		return fmt.Errorf("requested output must be in a different format than input file (both are %s)", inFmt)
	}

	// Build output file path
	outputPath, err := buildOutputPath(inputFile, outFmt, outputFile)
	if err != nil {
		return err
	}

	// Read and parse the input file
	data, err := readFile(inputFile, inFmt)
	if err != nil {
		return err
	}

	// Convert and write to file
	return writeOutput(data, outFmt, outputPath)
}
