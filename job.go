package gotodoit

import (
	"context"
	"encoding/json"
	"log"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/model"
	"github.com/achiku/gotodoit/service"
	"github.com/achiku/qg"
	"github.com/pkg/errors"
)

// JobApp job global values
type JobApp struct {
	BaseApp
	Logger *log.Logger
}

// Context returns context
func (app *JobApp) Context() context.Context {
	return context.Background()
}

// UpdateUserInfoArgs update todo etc args
type UpdateUserInfoArgs struct {
	UserID   string
	Email    string
	Username string
	Status   string
}

// UpdateUserInfo update user info job
func (app *JobApp) UpdateUserInfo(j *qg.Job) error {
	var args UpdateUserInfoArgs
	if err := json.Unmarshal(j.Args, &args); err != nil {
		return errors.Wrap(err, "failed to unmarshal UpdateTodoETCArgs")
	}
	dbReq := &model.TodoUser{
		UUID:     args.UserID,
		Username: args.Username,
		Email:    args.Email,
		Status:   args.Status,
	}
	if err := service.UpdateUser(j.Tx(), dbReq); err != nil {
		return errors.Wrap(err, "service.UpdateUser failed")
	}

	ctx := app.Context()
	req := &estc.UpdateUserRequest{
		ID:       args.UserID,
		UserName: args.Username,
		Email:    args.Email,
	}
	u, err := app.EstcClient.UpdateUser(ctx, req)
	if err != nil {
		return errors.Wrap(err, "EstcClient.UpdateUser failed")
	}
	app.Logger.Printf("userID=%s,email=%s", u.ID, u.Email)
	return nil
}
