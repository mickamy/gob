package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mickamy/gob"
)

type Postgres struct {
	cfg gob.Database
}

func (pg *Postgres) dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		pg.cfg.Host, pg.cfg.Port, pg.cfg.User, pg.cfg.Password)
}

func New(cfg gob.Database) *Postgres {
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
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return true, fmt.Errorf("failed to check if database exists: %w", err)
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
