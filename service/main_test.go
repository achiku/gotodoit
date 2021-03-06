package service

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/achiku/gotodoit/model"
)

func TestMain(m *testing.M) {
	flag.Parse()

	dbSetupCfg := model.DBConfig{
		DBName:  "gotodoit",
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
		User:    "gotodoit_root",
	}
	tblSetupCfg := model.DBConfig{
		DBName:  "gotodoit",
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
		User:    "gotodoit_api_test",
	}
	testSchema := "gotodoit_api_test"
	testUser := "gotodoit_api_test"

	model.TestDropSchema(dbSetupCfg, testSchema)

	if err := model.TestCreateSchema(dbSetupCfg, testSchema, testUser); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := model.TestCreateTables(tblSetupCfg, ".."); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	if err := model.TestDropSchema(dbSetupCfg, testSchema); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}
