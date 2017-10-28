package model

import "testing"

func TestUser_GetUserByAccessToken(t *testing.T) {
	tx, clean := TestSetupTx(t)
	defer clean()

	u := TestCreateUserData(t, tx, &TodoUser{})
	at := TestCreateAccessTokenData(t, tx, u)

	u, found, err := GetUserByAccessToken(tx, at.Token)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("must be found")
	}
	t.Logf("%v", u)
}

func TestUser_UpdateUser(t *testing.T) {
	tx, clean := TestSetupTx(t)
	defer clean()

	u := TestCreateUserData(t, tx, &TodoUser{})
	target := &TodoUser{
		UUID:     u.UUID,
		Username: "modified-username",
		Status:   "inactive",
		Email:    "modified-email@example.com",
	}
	if err := UpdateUser(tx, target); err != nil {
		t.Fatal(err)
	}

	res, err := GetTodoUserByPk(tx, u.UUID)
	if err != nil {
		t.Fatal(err)
	}
	if res.Status != target.Status {
		t.Errorf("want %s got %s", target.Status, res.Status)
	}
}
