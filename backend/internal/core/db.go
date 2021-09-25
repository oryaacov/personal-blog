package core

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"

	"github.com/oryaacov/personal-blog/internal/common"
)

type DB struct {
	DB   *sqlx.DB
	deps DBDeps
}

type DBDeps struct {
	fx.In
	Config common.Config
}

func InitDB(deps DBDeps) (*DB, error) {
	sqlxDB, err := sqlx.Connect(deps.Config.Database, deps.Config.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db with: %w", err)
	}

	err = sqlxDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db with: %w", err)
	}

	return &DB{DB: sqlxDB, deps: deps}, nil
}
