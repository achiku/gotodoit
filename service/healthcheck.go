package service

import "github.com/achiku/gotodoit/model"

// Healthcheck health check
func Healthcheck(tx model.Queryer) (string, error) {
	t, err := model.Healthcheck(tx)
	if err != nil {
		return "<`1`>", err
	}
	return t, nil
}
