package service

import (
	"context"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/model"
	"github.com/pkg/errors"
)

// Todo todo
type Todo struct {
	*model.Todo
	EstimatedTime int
}

// GetTodosByUserID get todos by user id
func GetTodosByUserID(
	ctx context.Context, db model.Queryer, c *estc.Client, userID string, completed bool) ([]Todo, error) {
	tds, err := model.GetTodosByUserID(db, userID, completed)
	if err != nil {
		return nil, errors.Wrap(err, "model.GetTodsByUserID failed")
	}
	var ttds []Todo
	for _, t := range tds {
		req := &estc.Task{
			Name:       t.Name,
			Difficulty: 10,
		}
		etc, err := c.EstimateTimeToComplete(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "c.EstimateTimeToComplete failed")
		}
		td := Todo{
			Todo:          &t,
			EstimatedTime: etc.Time,
		}
		ttds = append(ttds, td)
	}
	return ttds, nil
}

// GetUserTodoByID get todo by id
func GetUserTodoByID(
	ctx context.Context, db model.Queryer, c *estc.Client, userID, todoID string) (*Todo, bool, error) {
	td, found, err := model.GetUserTodoByID(db, userID, todoID)
	if err != nil {
		return nil, false, errors.Wrap(err, "model.GetUserTodoByID failed")
	}
	if !found {
		return nil, false, nil
	}
	req := &estc.Task{
		Name:       td.Name,
		Difficulty: 10,
	}
	etc, err := c.EstimateTimeToComplete(ctx, req)
	if err != nil {
		return nil, false, errors.Wrap(err, "c.EstimateTimeToComplete failed")
	}
	t := &Todo{
		Todo:          td,
		EstimatedTime: etc.Time,
	}
	return t, true, nil
}
