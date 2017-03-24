package model

import "database/sql"

// GetTodosByUserID get todos by user id
func GetTodosByUserID(tx Queryer, userID string, completed bool) ([]Todo, error) {
	rows, err := tx.Query(`
	select
		uuid
		, user_id
		, name
		, duration
		, started_at
		, is_completed
	from todo
	where user_id = $1
	and is_completed = $2
	order by uuid
	`, userID, completed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var td Todo
		err := rows.Scan(
			&td.UUID,
			&td.UserID,
			&td.Name,
			&td.Duration,
			&td.StartedAt,
			&td.IsCompleted,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, td)
	}
	return todos, nil
}

// GetUserTodoByID get todos by user id
func GetUserTodoByID(tx Queryer, userID, todoID string) (*Todo, bool, error) {
	var td Todo
	err := tx.QueryRow(`
	select
		uuid
		, user_id
		, name
		, duration
		, started_at
		, is_completed
	from todo
	where user_id = $1
	and uuid = $2
	`, userID, todoID).Scan(
		&td.UUID,
		&td.UserID,
		&td.Name,
		&td.Duration,
		&td.StartedAt,
		&td.IsCompleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &td, true, nil
}
