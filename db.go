package gotodoit

import (
	"database/sql"

	"github.com/achiku/gotodoit/model"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib" // pgx
)

// NewDB creates db
func NewDB(cfg *Config) (model.DBer, error) {
	dbCfg := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     cfg.DBHost,
			User:     cfg.DBUser,
			Password: cfg.DBPass,
			Database: cfg.DBName,
			Port:     cfg.DBPort,
		},
	}
	stdlib.RegisterDriverConfig(dbCfg)
	db, err := sql.Open("pgx", dbCfg.ConnectionString(""))
	if err != nil {
		return nil, err
	}
	return db, nil
}
