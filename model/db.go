package model

import (
	"database/sql"
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
