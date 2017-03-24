package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	txdb "github.com/achiku/pgtxdb"
)

func init() {
	txdb.Register("txdb", "postgres", "postgres://gotodoit_api_test@localhost:5432/gotodoit?sslmode=disable")
}

// ConnStr creates conn string from config
func ConnStr(c DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=%s", c.User, c.UserPass, c.Port, c.DBName, c.SSLMode)
}

// TestCreateSchema set up test schema
func TestCreateSchema(cfg DBConfig, schema, user string) error {
	db, err := sql.Open("postgres", ConnStr(cfg))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA %s AUTHORIZATION %s", schema, user))
	if err != nil {
		log.Printf("failed to create test schema: %s", schema)
		return err
	}
	return nil
}

// TestDropSchema set up test schema
func TestDropSchema(cfg DBConfig, schema string) error {
	db, err := sql.Open("postgres", ConnStr(cfg))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP SCHEMA %s CASCADE", schema))
	if err != nil {
		log.Printf("failed to create test schema: %s", schema)
		return err
	}
	return nil
}

// TestCreateTables create test tables
func TestCreateTables(cfg DBConfig, path string) error {
	orgPwd, _ := os.Getwd()
	defer func() {
		os.Chdir(orgPwd)
	}()

	os.Chdir(path)
	cmd := exec.Command("alembic", "upgrade", "head", "--sql")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("failed to execute alembic:\n %s", stderr.String())
		return err
	}

	db, err := sql.Open("postgres", ConnStr(cfg))
	if err != nil {
		return err
	}
	_, err = db.Exec(stdout.String())
	if err != nil {
		log.Println("failed to create test tables")
		return err
	}
	return nil
}

// TestSetupTx create tx and cleanup func for test
func TestSetupTx(t *testing.T) (Txer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		tx.Rollback()
		db.Close()
	}
	return tx, cleanup
}

// TestSetupDB create db and cleanup func for test
func TestSetupDB(t *testing.T) (DBer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		db.Close()
	}
	return db, cleanup
}
