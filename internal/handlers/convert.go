package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

// Convert reads a config file and writes it in the target format.
func Convert(inputPath, outFmt, outputDir string) (string, error) {
	inFmt, err := DetectFormat(inputPath)
	if err != nil {
		return "", err
	}

	if inFmt == outFmt {
		return "", fmt.Errorf("requested output must be in a different format than input file (both are %s)", inFmt)
	}

	data, err := ReadFile(inputPath, inFmt)
	if err != nil {
		return "", err
	}

	outputPath, err := buildOutputPath(inputPath, outFmt, outputDir)
	if err != nil {
		return "", err
	}

	if err := writeOutputFile(data, outFmt, outputPath); err != nil {
		return "", err
	}

	return outputPath, nil
}

func buildOutputPath(inputPath, outFmt, dir string) (string, error) {
	if dir == "" {
		dir = os.TempDir()
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory %q: %w", dir, err)
	}

	baseName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return filepath.Join(dir, baseName+"."+outFmt), nil
}

func writeOutputFile(data map[string]interface{}, format, outputPath string) error {
	var output []byte
	var err error

	switch format {
	case "json":
		output, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to convert to JSON: %w", err)
		}
		output = append(output, '\n')
	case "yaml":
		output, err = yaml.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to convert to YAML: %w", err)
		}
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
