package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // sql database
	"github.com/pkg/errors"
)

// DB database
type DB struct {
	*sql.DB
}

// DBConfig config
type DBConfig struct {
	Host     string
	User     string
	UserPass string
	Port     string
	DBName   string
	SSLMode  string
}

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// Txer database/sql transaction interface
type Txer interface {
	Queryer
	Commit() error
	Rollback() error
}

// DBer database/sql
type DBer interface {
	Queryer
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}

// NewDB creates DB
func NewDB(c *DBConfig) (DBer, error) {
	conStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", c.User, c.DBName, c.SSLMode)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db")
	}
	return &DB{db}, nil
}
