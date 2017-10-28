package gotodoit

import (
	"net/http"

	"github.com/achiku/gotodoit/iapi"
	"github.com/achiku/gotodoit/service"
	"github.com/achiku/mux"
	"github.com/pkg/errors"
)

// Healthcheck healthcheck
func (app *APIApp) Healthcheck(
	w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	m, err := service.Healthcheck(app.DB)
	if err != nil {
		c, e := iapi.NewInternalServerError()
		return c, e, errors.Wrap(err, "service.Healthcheck failed")
	}
	res := iapi.HealthcheckSelfResponse{
		Message: m,
	}
	return http.StatusOK, res, nil
}

// GetUserDetail get user
func (app *APIApp) GetUserDetail(
	w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	auth := getAuthData(r.Context())
	u, err := service.GetUserByID(app.DB, auth.User.UUID)
	if err != nil {
		c, e := iapi.NewInternalServerError()
		return c, e, errors.Wrap(err, "service.GetUserByID failed")
	}
	res := iapi.UserSelfResponse{
		ID:       u.UUID,
		Username: u.Username,
		Email:    u.Email,
	}
	return http.StatusOK, res, nil
}

// GetTodos get todos
func (app *APIApp) GetTodos(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	auth := getAuthData(r.Context())
	tds, err := service.GetTodosByUserID(r.Context(), app.DB, app.EstcClient, auth.User.UUID, false)
	if err != nil {
		c, e := iapi.NewInternalServerError()
		return c, e, errors.Wrap(err, "service.GetTodosByUserID failed")
	}
	var res iapi.TodoInstancesResponse
	for _, t := range tds {
		res = append(res, iapi.Todo{
			ID:            t.UUID,
			Name:          t.Name,
			TotalDuration: t.Duration,
			StartedAt:     t.StartedAt,
		})
	}
	return http.StatusOK, res, nil
}

// GetTodoByID get todo by id
func (app *APIApp) GetTodoByID(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	todoID := mux.Vars(r)["todoID"]
	auth := getAuthData(r.Context())
	td, found, err := service.GetUserEstimatedTodoByID(
		r.Context(), app.DB, app.EstcClient, auth.User.UUID, todoID)
	if err != nil {
		c, e := iapi.NewInternalServerError()
		return c, e, errors.Wrap(err, "service.GetTodoByID failed")
	}
	if !found {
		c, e := iapi.NewNotFoundError()
		return c, e, errors.Errorf("todoID=%s not found", todoID)
	}
	res := iapi.TodoSelfResponse{
		ID:            td.UUID,
		Name:          td.Name,
		StartedAt:     td.StartedAt,
		TotalDuration: td.Duration,
	}
	return http.StatusOK, res, nil
}
