# goko-cli

A small command-line tool written in Go that converts configuration files between **JSON** and **YAML**.

Built with [Cobra](https://github.com/spf13/cobra) and [goccy/go-yaml](https://github.com/goccy/go-yaml).

## Features

- Convert `.json` → `.yaml`
- Convert `.yaml` / `.yml` → `.json`
- Auto-detects the input format from the file extension
- Writes output to a directory you choose, or falls back to the system temp directory
- Friendly errors for unsupported extensions, missing files, and same-format conversions

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/ridwanulhoquejr/goko-cli.git
cd goko-cli
go build -o goku ./cmd
```

Or install directly with `go install`:

```bash
go install github.com/ridwanulhoquejr/goko-cli/cmd@latest
```

> Requires Go **1.24** or newer.

## Usage

```
goku -i <file_path> -o <json|yaml> [-d <output_dir>]
```

### Flags

| Flag           | Short | Description                                            |
| -------------- | ----- | ------------------------------------------------------ |
| `--input`      | `-i`  | Path to the input file (`.json`, `.yaml`, or `.yml`)   |
| `--output`     | `-o`  | Target format: `json` or `yaml`                        |
| `--dir`        | `-d`  | Output directory (defaults to the system temp dir)     |

### Examples

Convert a JSON file to YAML and place the result in `./output`:

```bash
goku -i example.json -o yaml -d ./output
```

Convert a YAML file to JSON using the default output location:

```bash
goku -i config.yaml -o json
```

### Example

Given `example.json`:

```json
{
    "Name": "Goku CLI",
    "Version": "1.0.0",
    "Description": "A command-line interface for Goku, the powerful AI assistant."
}
```

Running `goku -i example.json -o yaml -d ./output` produces `output/example.yaml`:

```yaml
Description: A command-line interface for Goku, the powerful AI assistant.
Name: Goku CLI
Version: 1.0.0
```

## Project Structure

```
.
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   └── command/
│       ├── command.go       # Cobra root command & flags
│       └── convert.go       # Read, convert, and write logic
├── example.json             # Sample input
├── output/                  # Sample output
├── go.mod
└── go.sum
```

## License

MIT
