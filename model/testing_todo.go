package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/achiku/hseq"
	"github.com/achiku/ovw"
)

// TestCreateTodoData create todo test data
func TestCreateTodoData(t *testing.T, tx Queryer, u *TodoUser, td *Todo) *Todo {
	n := time.Now()
	tdDefault := Todo{
		Duration:    100,
		Name:        fmt.Sprintf("todo name %s", hseq.Get("todo.name")),
		IsCompleted: false,
		StartedAt:   n.AddDate(0, 0, -1),
		UserID:      u.UUID,
	}
	var target Todo
	if err := ovw.MergeOverwrite(tdDefault, td, &target); err != nil {
		t.Fatal(err)
	}
	if err := target.Create(tx); err != nil {
		t.Fatal(err)
	}
	return &target
}
