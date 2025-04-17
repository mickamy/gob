package database

import (
	"fmt"

	"github.com/mickamy/gob"
	"github.com/mickamy/gob/internal/database/mysql"
	"github.com/mickamy/gob/internal/database/postgres"
)

func New(cfg gob.Database) (Database, error) {
	switch cfg.Driver {
	case "postgres":
		return postgres.New(cfg), nil
	case "mysql":
		return mysql.New(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
}

var (
	_ Database = (*mysql.MySQL)(nil)
	_ Database = (*postgres.Postgres)(nil)
)
