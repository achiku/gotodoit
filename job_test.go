package gotodoit

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/model"
	"github.com/achiku/qg"
)

func TestUpdateUserInfo(t *testing.T) {
	app, tx, _, cleanup := testSetupJobApp(t)
	defer cleanup()

	u := model.TestCreateUserData(t, tx, &model.TodoUser{})
	args := UpdateUserInfoArgs{
		UserID:   u.UUID,
		Email:    fmt.Sprintf("%s@example.com", u.UUID),
		Status:   "inactive",
		Username: "updated-username",
	}
	jsonArgs, err := json.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(estc.TestNewMux(estc.DefaultHandlerMap))
	defer ts.Close()
	app.Config.EstcConfig.BaseEndpoint = ts.URL

	j := qg.TestInjectJobTx(&qg.Job{Args: jsonArgs}, tx)
	if err := app.UpdateUserInfo(j); err != nil {
		t.Fatal(err)
	}
}
