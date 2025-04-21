package godb

import (
	"fmt"

	"github.com/mickamy/godb/config"
)

func Reset(cfg config.Config) error {
	if err := Drop(cfg); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	if err := Create(cfg); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	if err := Migrate(cfg); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
