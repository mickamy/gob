package database

import (
	"fmt"

	"github.com/mickamy/godb/config"
	"github.com/mickamy/godb/internal/database/mysql"
	"github.com/mickamy/godb/internal/database/postgres"
)

func New(cfg config.Database) (Database, error) {
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
