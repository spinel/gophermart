package bun

import (
	"database/sql"

	"github.com/spinel/gophermart/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*bun.DB
}

const notDeleted = "deleted_at is null"

// Dial creates new database connection to postgres
func Dial(cfg *config.Config) (*DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PgURL)))
	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
