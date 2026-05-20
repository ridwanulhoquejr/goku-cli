package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

// DetectFormat determines the file format from its extension.
func DetectFormat(filePath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".json":
		return "json", nil
	case ".yaml", ".yml":
		return "yaml", nil
	default:
		return "", fmt.Errorf("unsupported file extension %q (supported: .json, .yaml, .yml)", ext)
	}
}

// ReadFile reads a file and unmarshals it into a generic map.
func ReadFile(filePath, format string) (map[string]interface{}, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var result map[string]interface{}

	switch format {
	case "json":
		if err := json.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	case "yaml":
		if err := yaml.Unmarshal(raw, &result); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	}

	return result, nil
}

// DeriveResourceName extracts the filename without extension.
func DeriveResourceName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}
