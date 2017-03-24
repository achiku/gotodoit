package model

import "testing"

func TestNewDB(t *testing.T) {
	c := &DBConfig{
		Host:    "localhost",
		User:    "pgtest",
		DBName:  "pgtest",
		SSLMode: "disable",
		Port:    "5432",
	}
	db, err := NewDB(c)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
