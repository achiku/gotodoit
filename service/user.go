package service

import (
	"github.com/achiku/gotodoit/model"
	"github.com/pkg/errors"
)

// GetUserByID get user by id
func GetUserByID(tx model.Queryer, userID string) (*model.TodoUser, error) {
	u, err := model.GetTodoUserByPk(tx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "model.GetTodoByPk failed")
	}
	return u, nil
}

// GetUserByAccessToken get user by access token
func GetUserByAccessToken(tx model.Queryer, token string) (*model.TodoUser, bool, error) {
	u, found, err := model.GetUserByAccessToken(tx, token)
	if err != nil {
		return nil, false, errors.Wrap(err, "model.GetUserByAccessToken failed")
	}
	if !found {
		return nil, false, nil
	}
	return u, true, nil
}

// CreateUser create user
func CreateUser(tx model.Queryer, u *model.TodoUser) (*model.TodoUser, error) {
	u, err := model.CreateUser(tx, u)
	if err != nil {
		return nil, errors.Wrap(err, "model.CreateUser failed")
	}
	return u, nil
}
