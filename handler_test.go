package gotodoit

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/iapi"
	"github.com/achiku/gotodoit/model"
)

func TestHealthCheck(t *testing.T) {
	app, _, _, cleanup := testSetupApp(t)
	defer cleanup()

	req := testCreateRequest(t, "GET", "/v1/healthcheck", nil)
	wr := httptest.NewRecorder()
	status, res, err := app.Healthcheck(wr, req)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusOK {
		t.Errorf("want status=%d got %d", http.StatusOK, status)
	}
	hc, ok := res.(iapi.HealthcheckSelfResponse)
	if !ok {
		t.Errorf("want iapi.HealthcheckSelfResponse")
	}
	if hc.Message == "" {
		t.Errorf("want hc.Message not blank")
	}
}

func TestGetTodos(t *testing.T) {
	app, tx, ctx, cleanup := testSetupApp(t)
	defer cleanup()

	u := model.TestCreateUserData(t, tx, &model.TodoUser{})
	model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name:        "buy milk",
		IsCompleted: false,
	})
	model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name:        "done todo",
		IsCompleted: false,
	})

	ts := httptest.NewServer(estc.TestNewMux(estc.DefaultHandlerMap))
	defer ts.Close()

	app.EstcClient = estc.NewClient(estc.TestNewConfig(ts.URL), &http.Client{}, nil)
	req := testCreateRequest(t, "GET", "/v1/todos", nil).WithContext(
		context.WithValue(ctx, ctxKeyAuth, AuthModel{Token: "token", User: u}),
	)
	wr := httptest.NewRecorder()
	status, res, err := app.GetTodos(wr, req)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusOK {
		t.Errorf("want status=%d got %d", http.StatusOK, status)
	}
	td, ok := res.(iapi.TodoInstancesResponse)
	if !ok {
		t.Errorf("want iapi.TodoSelfResponse")
	}
	if expected := 2; len(td) != expected {
		t.Errorf("want len(td)=%d got %d", expected, len(td))
	}
}
