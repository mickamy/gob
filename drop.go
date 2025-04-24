package godb

import (
	"fmt"

	"github.com/mickamy/godb/config"
	"github.com/mickamy/godb/internal/database"
)

func Drop(cfg config.Config, force bool) error {
	db, err := database.New(cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	if err := db.Drop(force); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	return nil
}
