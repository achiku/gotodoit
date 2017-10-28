package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/model"
)

func TestTodo_GetTodosByUserID(t *testing.T) {
	tx, clean := model.TestSetupTx(t)
	defer clean()

	u := model.TestCreateUserData(t, tx, &model.TodoUser{})
	notCompleted := model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name:        "buy milk",
		IsCompleted: false,
	})
	completed := model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name:        "done todo",
		IsCompleted: true,
	})

	data := []struct {
		Todo *model.Todo
		Done bool
	}{
		{Todo: notCompleted, Done: false},
		{Todo: completed, Done: true},
	}
	ts := httptest.NewServer(estc.TestNewMux(estc.DefaultHandlerMap))
	defer ts.Close()

	client := estc.NewClient(estc.TestNewConfig(ts.URL), &http.Client{}, nil)
	ctx := context.Background()
	for _, d := range data {
		tds, err := GetTodosByUserID(ctx, tx, client, u.UUID, d.Done)
		if err != nil {
			t.Fatal(err)
		}
		if result := len(tds); result != 1 {
			t.Errorf("want 1 got %d", result)
		}
		if result := tds[0]; result.Name != d.Todo.Name {
			t.Errorf("want %s got %s", d.Todo.Name, result.Name)
		}
	}
}

func TestTodo_GetUserTodoByID(t *testing.T) {
	tx, clean := model.TestSetupTx(t)
	defer clean()

	u := model.TestCreateUserData(t, tx, &model.TodoUser{})
	td1 := model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name: "buy milk",
	})
	td2 := model.TestCreateTodoData(t, tx, u, &model.Todo{
		Name: "done todo",
	})

	data := []struct {
		Todo  *model.Todo
		Found bool
		ID    string
	}{
		{Todo: td1, ID: td1.UUID, Found: true},
		{Todo: td2, ID: td2.UUID, Found: true},
	}
	ts := httptest.NewServer(estc.TestNewMux(estc.DefaultHandlerMap))
	defer ts.Close()

	client := estc.NewClient(estc.TestNewConfig(ts.URL), &http.Client{}, nil)
	ctx := context.Background()
	for _, d := range data {
		td, found, err := GetUserEstimatedTodoByID(ctx, tx, client, u.UUID, d.ID)
		if err != nil {
			t.Fatal(err)
		}
		if found != d.Found {
			t.Errorf("want found=%t got %t", d.Found, found)
		}
		if td.Name != d.Todo.Name {
			t.Errorf("want Name=%s got %s", d.Todo.Name, td.Name)
		}
	}
}
