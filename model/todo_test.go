package model

import "testing"

func TestTodo_GetTodosByUserID(t *testing.T) {
	tx, clean := TestSetupTx(t)
	defer clean()

	u := TestCreateUserData(t, tx, &TodoUser{})
	notCompleted := TestCreateTodoData(t, tx, u, &Todo{
		Name:        "buy milk",
		IsCompleted: false,
	})
	completed := TestCreateTodoData(t, tx, u, &Todo{
		Name:        "done todo",
		IsCompleted: true,
	})

	data := []struct {
		Todo *Todo
		Done bool
	}{
		{Todo: notCompleted, Done: false},
		{Todo: completed, Done: true},
	}

	for _, d := range data {
		tds, err := GetTodosByUserID(tx, u.UUID, d.Done)
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
	tx, clean := TestSetupTx(t)
	defer clean()

	u := TestCreateUserData(t, tx, &TodoUser{})
	td1 := TestCreateTodoData(t, tx, u, &Todo{
		Name: "buy milk",
	})
	td2 := TestCreateTodoData(t, tx, u, &Todo{
		Name: "done todo",
	})

	data := []struct {
		Todo *Todo
		ID   string
	}{
		{Todo: td1, ID: td1.UUID},
		{Todo: td2, ID: td2.UUID},
	}

	for _, d := range data {
		td, found, err := GetUserTodoByID(tx, u.UUID, d.ID)
		if err != nil {
			t.Fatal(err)
		}
		if !found {
			t.Error("need to be found")
		}
		if td.Name != d.Todo.Name {
			t.Errorf("want %s got %s", d.Todo.Name, td.Name)
		}
	}
}
