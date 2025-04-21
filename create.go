package godb

import (
	"errors"
	"fmt"

	"github.com/mickamy/godb/config"
	"github.com/mickamy/godb/internal/database"
)

var (
	ErrCreateDatabaseExists = errors.New("database already exists")
)

func Create(cfg config.Config) error {
	db, err := database.New(cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	exists, err := db.Exists()
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}
	if exists {
		return ErrCreateDatabaseExists
	}

	if err := db.Create(); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	return nil
}
