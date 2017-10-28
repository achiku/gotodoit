package iapi

import "time"

// Healthcheck struct for healthcheck resource
type Healthcheck struct {
	Message string `json:"message"`
}

// Todo struct for todo resource
type Todo struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	StartedAt     time.Time `json:"startedAt"`
	StoppedAt     time.Time `json:"stoppedAt,omitempty"`
	TotalDuration int64     `json:"totalDuration"`
}

// User struct for user resource
type User struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	Username string `json:"username"`
}

// HealthcheckSelfResponse struct for healthcheck
// GET: /healthcheck
type HealthcheckSelfResponse Healthcheck

// TodoInstancesRequest struct for todo
// GET: /todos
type TodoInstancesRequest struct {
	Limit  int64 `json:"limit,omitempty" schema:"limit"`
	Offset int64 `json:"offset,omitempty" schema:"offset"`
}

// TodoInstancesResponse struct for todo
// GET: /todos
type TodoInstancesResponse []Todo

// TodoSelfResponse struct for todo
// GET: /todos/{(#/definitions/todo/definitions/identity)}
type TodoSelfResponse Todo

// TodoCreateRequest struct for todo
// POST: /todos
type TodoCreateRequest struct {
	Name   string `json:"name"`
	UserID string `json:"userId,omitempty"`
}

// TodoCreateResponse struct for todo
// POST: /todos
type TodoCreateResponse Todo

// UserSelfResponse struct for user
// GET: /users/me
type UserSelfResponse User

// UserCreateRequest struct for user
// POST: /users
type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Username string `json:"username"`
}

// UserCreateResponse struct for user
// POST: /users
type UserCreateResponse struct {
	Token string `json:"token,omitempty"`
	User  *User  `json:"user,omitempty"`
}
