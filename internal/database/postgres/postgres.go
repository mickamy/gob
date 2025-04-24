package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mickamy/godb/config"
)

type Postgres struct {
	cfg config.Database
}

func (pg *Postgres) dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		pg.cfg.Host, pg.cfg.Port, pg.cfg.User, pg.cfg.Password)
}

func New(cfg config.Database) *Postgres {
	return &Postgres{cfg: cfg}
}

func (pg *Postgres) Name() string {
	return pg.cfg.Name
}

func (pg *Postgres) Exists() (bool, error) {
	db, err := sql.Open("postgres", pg.dsn())
	if err != nil {
		return false, fmt.Errorf("failed to open database connection: %w", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	var exists bool
	query := `SELECT 1 FROM pg_database WHERE datname = $1`
	err = db.QueryRow(query, pg.cfg.Name).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if database exists: %w", err)
	}
	return exists, nil
}

func (pg *Postgres) Create() error {
	db, err := sql.Open("postgres", pg.dsn())
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec("CREATE DATABASE " + pg.cfg.Name)
	return err
}

func (pg *Postgres) Drop(force bool) error {
	db, err := sql.Open("postgres", pg.dsn())
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", pg.cfg.Name)
	if force {
		query += " WITH (FORCE)"
	}
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	return nil
}
