package service

import (
	"testing"

	"github.com/achiku/gotodoit/model"
)

func TestUser_GetUserByID(t *testing.T) {
	tx, cleanup := model.TestSetupTx(t)
	defer cleanup()

	u := model.TestCreateUserData(t, tx, &model.TodoUser{
		Email: "akira.chiku@gmail.com",
	})
	tu, err := GetUserByID(tx, u.UUID)
	if err != nil {
		t.Fatal(err)
	}
	if tu.Email != u.Email {
		t.Errorf("want %s got %s", u.Email, tu.Email)
	}
}
