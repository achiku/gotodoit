package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/achiku/hseq"
	"github.com/achiku/ovw"
)

// TestCreateUserData create user test data
func TestCreateUserData(t *testing.T, tx Queryer, u *TodoUser) *TodoUser {
	uDefault := TodoUser{
		Email: fmt.Sprintf(
			"akira.chiku.%s@gmail.com",
			hseq.Get("todo_user.email")),
		Username: fmt.Sprintf(
			"user%s",
			hseq.Get("todo_user.username")),
		Status: "active",
	}
	var target TodoUser
	if err := ovw.MergeOverwrite(uDefault, u, &target); err != nil {
		t.Fatal(err)
	}
	if err := target.Create(tx); err != nil {
		t.Fatal(err)
	}
	return &target
}

// TestCreateAccessTokenData create test data
func TestCreateAccessTokenData(t *testing.T, tx Queryer, u *TodoUser) *AccessToken {
	at := AccessToken{
		GeneratedAt: time.Now(),
		IsActive:    true,
		Token:       fmt.Sprintf("token%s", hseq.Get("access_token.token")),
		UserID:      u.UUID,
	}
	if err := at.Create(tx); err != nil {
		t.Fatal(err)
	}
	return &at
}
