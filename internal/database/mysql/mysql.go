package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mickamy/gob/config"
)

type MySQL struct {
	cfg config.Database
}

func (m *MySQL) dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		m.cfg.User,
		m.cfg.Password,
		m.cfg.Host,
		m.cfg.Port,
		dbname,
	)
}

func New(cfg config.Database) *MySQL {
	return &MySQL{cfg: cfg}
}

func (m *MySQL) Name() string {
	return m.cfg.Name
}

func (m *MySQL) Exists() (bool, error) {
	db, err := sql.Open("mysql", m.dsn("information_schema"))
	if err != nil {
		return false, err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	var exists string
	query := "SELECT SCHEMA_NAME FROM SCHEMATA WHERE SCHEMA_NAME = ?"
	err = db.QueryRow(query, m.cfg.Name).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if database exists: %w", err)
	}
	return exists != "", err
}

func (m *MySQL) Create() error {
	db, err := sql.Open("mysql", m.dsn(""))
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", m.cfg.Name))
	return err
}

func (m *MySQL) Drop() error {
	db, err := sql.Open("mysql", m.dsn(m.cfg.Name))
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", m.cfg.Name))
	if err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	return nil
}
