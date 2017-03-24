package iapi

import (
	"time"
)

// Healthcheck healthcheck resource
type Healthcheck struct {
	Message string `json:"message" schema:"message"`
}

// Todo todo resource
type Todo struct {
	ID            string    `json:"id" schema:"id"`
	Name          string    `json:"name" schema:"name"`
	StartedAt     time.Time `json:"startedAt" schema:"startedAt"`
	StoppedAt     time.Time `json:"stoppedAt,omitempty" schema:"stoppedAt"`
	TotalDuration int64     `json:"totalDuration" schema:"totalDuration"`
}

// User user resource
type User struct {
	Email    string `json:"email" schema:"email"`
	ID       string `json:"id" schema:"id"`
	Username string `json:"username" schema:"username"`
}

// HealthcheckSelfResponse response
type HealthcheckSelfResponse Healthcheck

// TodoCreateRequest request
type TodoCreateRequest struct {
	Name   string `json:"name" schema:"name"`
	UserID string `json:"userId" schema:"userId"`
}

// TodoCreateResponse response
type TodoCreateResponse Todo

// TodoInstancesRequest request
type TodoInstancesRequest struct {
	Limit  int64 `json:"limit,omitempty" schema:"limit"`
	Offset int64 `json:"offset,omitempty" schema:"offset"`
}

// TodoInstancesResponse response
type TodoInstancesResponse []struct {
	ID            string    `json:"id" schema:"id"`
	Name          string    `json:"name" schema:"name"`
	StartedAt     time.Time `json:"startedAt" schema:"startedAt"`
	StoppedAt     time.Time `json:"stoppedAt,omitempty" schema:"stoppedAt"`
	TotalDuration int64     `json:"totalDuration" schema:"totalDuration"`
}

// TodoSelfResponse response
type TodoSelfResponse Todo

// UserCreateRequest request
type UserCreateRequest struct {
	Email    string `json:"email" schema:"email"`
	Password string `json:"password" schema:"password"`
	Username string `json:"username" schema:"username"`
}

// UserCreateResponse response
type UserCreateResponse struct {
	Token string `json:"token,omitempty" schema:"token"`
	User  *User  `json:"user,omitempty" schema:"user"`
}

// UserSelfResponse response
type UserSelfResponse User
