package model

import (
	"database/sql"
	"encoding/hex"
	"time"

	uuid "github.com/satori/go.uuid"
)

// User status
const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
)

// CreateUser creates user
func CreateUser(tx Queryer, u *TodoUser) (*TodoUser, error) {
	if err := u.Create(tx); err != nil {
		return nil, err
	}
	tkn := AccessToken{
		UserID:      u.UUID,
		GeneratedAt: time.Now(),
		IsActive:    true,
		Token:       hex.EncodeToString(uuid.NewV4().Bytes()),
	}
	if err := tkn.Create(tx); err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserByAccessToken get user by access token
func GetUserByAccessToken(tx Queryer, token string) (*TodoUser, bool, error) {
	var u TodoUser
	err := tx.QueryRow(`
	select
		u.uuid
		,u.username
		,u.email
		,u.status
	from todo_user u
	join access_token a on u.uuid = a.user_id
	where a.token = $1
	and a.is_active = true
	`, token).Scan(
		&u.UUID,
		&u.Username,
		&u.Email,
		&u.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &u, true, nil
}
