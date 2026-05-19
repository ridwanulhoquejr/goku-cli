package db

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const defaultDSN = "postgres://postgres:postgres@localhost:5432/goku?sslmode=disable"

func GetDSN() string {
	dsn := os.Getenv("GOKU_DB_URL")
	if dsn == "" {
		dsn = defaultDSN
	}
	return dsn
}

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

var (
	instance *sqlx.DB
	once     sync.Once
	initErr  error
)

func Get() (*sqlx.DB, error) {
	once.Do(func() {
		instance, initErr = Connect()
	})
	return instance, initErr
}

func Close() error {
	if instance == nil {
		return nil
	}
	return instance.Close()
}

func RunMigrations(path string) error {
	m, err := migrate.New(path, GetDSN())
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

func RollbackMigrations(path string) error {
	m, err := migrate.New(path, GetDSN())
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}
	return nil
}
