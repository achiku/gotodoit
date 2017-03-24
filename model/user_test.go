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
