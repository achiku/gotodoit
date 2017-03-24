package model

import "time"

// Healthcheck health check
func Healthcheck(tx Queryer) (string, error) {
	var t time.Time
	err := tx.QueryRow(`
	select now()
	`).Scan(&t)
	if err != nil {
		return "<`1`>", err
	}
	return t.String(), nil
}
