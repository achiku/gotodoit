package model

import (
	"time"

	"github.com/pkg/errors"
)

// AccessToken represents gotodoit_api.access_token
type AccessToken struct {
	Token       string    // token
	UserID      string    // user_id
	GeneratedAt time.Time // generated_at
	IsActive    bool      // is_active
}

// Create inserts the AccessToken to the database.
func (r *AccessToken) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO access_token (token, user_id, generated_at, is_active) VALUES ($1, $2, $3, $4)`,
		&r.Token, &r.UserID, &r.GeneratedAt, &r.IsActive)
	if err != nil {
		return errors.Wrap(err, "failed to insert access_token")
	}
	return nil
}

// GetAccessTokenByPk select the AccessToken from the database.
func GetAccessTokenByPk(db Queryer, pk0 string) (*AccessToken, error) {
	var r AccessToken
	err := db.QueryRow(
		`SELECT token, user_id, generated_at, is_active FROM access_token WHERE token = $1`,
		pk0).Scan(&r.Token, &r.UserID, &r.GeneratedAt, &r.IsActive)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select access_token")
	}
	return &r, nil
}

// Todo represents gotodoit_api.todo
type Todo struct {
	UUID        string    // uuid
	UserID      string    // user_id
	Name        string    // name
	Duration    int64     // duration
	StartedAt   time.Time // started_at
	IsCompleted bool      // is_completed
}

// Create inserts the Todo to the database.
func (r *Todo) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO todo (user_id, name, duration, started_at, is_completed) VALUES ($1, $2, $3, $4, $5) RETURNING uuid`,
		&r.UserID, &r.Name, &r.Duration, &r.StartedAt, &r.IsCompleted).Scan(&r.UUID)
	if err != nil {
		return errors.Wrap(err, "failed to insert todo")
	}
	return nil
}

// GetTodoByPk select the Todo from the database.
func GetTodoByPk(db Queryer, pk0 string) (*Todo, error) {
	var r Todo
	err := db.QueryRow(
		`SELECT uuid, user_id, name, duration, started_at, is_completed FROM todo WHERE uuid = $1`,
		pk0).Scan(&r.UUID, &r.UserID, &r.Name, &r.Duration, &r.StartedAt, &r.IsCompleted)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select todo")
	}
	return &r, nil
}

// TodoUser represents gotodoit_api.todo_user
type TodoUser struct {
	UUID     string // uuid
	Username string // username
	Email    string // email
	Password string // password
	Status   string // status
}

// Create inserts the TodoUser to the database.
func (r *TodoUser) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO todo_user (username, email, password, status) VALUES ($1, $2, $3, $4) RETURNING uuid`,
		&r.Username, &r.Email, &r.Password, &r.Status).Scan(&r.UUID)
	if err != nil {
		return errors.Wrap(err, "failed to insert todo_user")
	}
	return nil
}

// GetTodoUserByPk select the TodoUser from the database.
func GetTodoUserByPk(db Queryer, pk0 string) (*TodoUser, error) {
	var r TodoUser
	err := db.QueryRow(
		`SELECT uuid, username, email, password, status FROM todo_user WHERE uuid = $1`,
		pk0).Scan(&r.UUID, &r.Username, &r.Email, &r.Password, &r.Status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select todo_user")
	}
	return &r, nil
}
