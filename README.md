# Goku CLI

A command-line toolkit written in Go for working with configuration files and managing them as resources in a PostgreSQL database.

Built with [Cobra](https://github.com/spf13/cobra), [sqlx](https://github.com/jmoiron/sqlx), [golang-migrate](https://github.com/golang-migrate/migrate), and [goccy/go-yaml](https://github.com/goccy/go-yaml).

## Features

- **Convert** between `JSON` and `YAML` (auto-detects input format from extension)
- **Persist** JSON/YAML resources in PostgreSQL as `JSONB`
- **CRUD** over stored resources: `save`, `list`, `get`, `update`, `delete`
- **Migrations** via `goku migrate up` / `goku migrate down`
- Configurable database connection through the `GOKU_DB_URL` environment variable

## Requirements

- Go **1.24** or newer
- PostgreSQL (for any command that touches the database)

## Installation

```bash
git clone https://github.com/ridwanulhoquejr/goku-cli.git
cd goku-cli
go build -o goku ./cmd
```

Or install directly:

```bash
go install github.com/ridwanulhoquejr/goku-cli/cmd@latest
```

## Configuration

The database connection string is read from `GOKU_DB_URL`. If unset, it defaults to:

```
postgres://postgres:postgres@localhost:5432/goku?sslmode=disable
```

Example `.env`:

```bash
GOKU_DB_URL=postgres://postgres:postgres@localhost:5432/goku?sslmode=disable
```

## Database setup

Apply the schema before using the resource commands:

```bash
goku migrate up
```

Roll back the most recent migration:

```bash
goku migrate down
```

The migrations directory defaults to `file://migrations`. Override with `--path` if needed.

## Commands

### `convert` — JSON ⇄ YAML

```bash
goku convert -i <file> -o <json|yaml> [-d <output_dir>]
```

| Flag       | Short | Description                                          |
| ---------- | ----- | ---------------------------------------------------- |
| `--input`  | `-i`  | Path to input file (`.json`, `.yaml`, or `.yml`)     |
| `--output` | `-o`  | Target format: `json` or `yaml`                      |
| `--dir`    | `-d`  | Output directory (defaults to the system temp dir)   |

Example:

```bash
goku convert -i example.json -o yaml -d ./output
```

### `save` — store a resource

```bash
goku save -i <resource.json|resource.yaml>
```

Reads the file and inserts it into `resource_table` as a JSONB document.

### `list` — list all resources

```bash
goku list
```

Prints a table of `ID`, `NAME`, `TYPE`, and `CREATED AT`.

### `get` — fetch a resource by ID

```bash
goku get --id <id>
```

### `update` — replace a resource by ID

```bash
goku update --id <id> -i <new_file.json|new_file.yaml>
```

### `delete` — remove a resource by ID

```bash
goku delete --id <id>
```

### `migrate` — schema management

```bash
goku migrate up
goku migrate down
goku migrate up --path file://migrations
```

## Schema

`migrations/000001_create_resource_table.up.sql`:

```sql
CREATE TABLE IF NOT EXISTS resource_table (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    type        TEXT NOT NULL,
    data        JSONB NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);
```

## Example

Given `example.json`:

```json
{
    "Name": "Goku CLI",
    "Version": "1.0.0",
    "Phase": "02",
    "Description": "A command-line interface for Goku, the powerful AI assistant."
}
```

Convert to YAML:

```bash
goku convert -i example.json -o yaml -d ./output
```

Save into the database:

```bash
goku save -i example.json
```

## Project Structure

```
.
├── cmd/
│   └── main.go                       # Entry point
├── internal/
│   ├── command/                      # Cobra commands
│   │   ├── root.go
│   │   ├── convert.go
│   │   ├── save.go
│   │   ├── list.go
│   │   ├── get.go
│   │   ├── update.go
│   │   ├── delete.go
│   │   └── migrate.go
│   ├── handlers/                     # Business logic
│   │   ├── convert.go
│   │   ├── save.go
│   │   ├── list.go
│   │   ├── get.go
│   │   ├── update.go
│   │   ├── delete.go
│   │   └── helpers.go
│   └── db/                           # Database access & migrations
│       ├── db.go
│       └── resource.go
├── migrations/                       # SQL migration files
├── example.json
├── output/
├── go.mod
└── go.sum
```

## License

MIT
