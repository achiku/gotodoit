package model

import (
	"flag"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()

	dbSetupCfg := DBConfig{
		DBName:  "gotodoit",
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
		User:    "gotodoit_root",
	}
	tblSetupCfg := DBConfig{
		DBName:  "gotodoit",
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
		User:    "gotodoit_api_test",
	}
	testSchema := "gotodoit_api_test"
	testUser := "gotodoit_api_test"

	TestDropSchema(dbSetupCfg, testSchema)

	if err := TestCreateSchema(dbSetupCfg, testSchema, testUser); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := TestCreateTables(tblSetupCfg, ".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	if err := TestDropSchema(dbSetupCfg, testSchema); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}
