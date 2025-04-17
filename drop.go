package gob

import (
	"fmt"

	"github.com/mickamy/gob/config"
	"github.com/mickamy/gob/internal/database"
)

func Drop(cfg config.Config) error {
	db, err := database.New(cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	if err := db.Drop(); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	return nil
}
