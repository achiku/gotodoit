package gotodoit

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/achiku/gotodoit/estc"
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
	if err := model.TestCreateTables(tblSetupCfg, "."); err != nil {
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

func testSetupApp(t *testing.T) (*APIApp, model.DBer, context.Context, func()) {
	config, err := NewConfig("./conf/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	db, cleanup := model.TestSetupDB(t)

	app := &APIApp{
		BaseApp: BaseApp{
			Config: config,
		},
		DB: db,
	}
	ctx := context.Background()
	return app, db, ctx, func() {
		cleanup()
	}
}

func testSetupJobApp(t *testing.T) (*JobApp, model.Txer, context.Context, func()) {
	config, err := NewConfig("./conf/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	db, cleanup := model.TestSetupDB(t)

	client := estc.NewClient(
		config.EstcConfig, http.DefaultClient, log.New(os.Stdout, "[estc]:", log.Lshortfile))
	app := &JobApp{
		BaseApp: BaseApp{
			Config:     config,
			EstcClient: client,
		},
		Logger: log.New(os.Stdout, "[job]:", log.Lshortfile),
	}
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	return app, tx, ctx, func() {
		tx.Rollback()
		cleanup()
	}
}

func testCreateRequest(t *testing.T, method, path string, b io.Reader) *http.Request {
	req, err := http.NewRequest(method, path, b)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
