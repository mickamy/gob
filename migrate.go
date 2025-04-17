package gob

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/mickamy/gob/config"
)

var (
	ErrMigrateNoChange = errors.New("no change")
)

func Migrate(cfg config.Config) error {
	dbURL, err := cfg.Database.URL()
	if err != nil {
		return fmt.Errorf("failed to get database URL: %w", err)
	}

	m, err := migrate.New(fmt.Sprintf("file://%s", cfg.Migrations.Dir), dbURL)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return ErrMigrateNoChange
		}
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
