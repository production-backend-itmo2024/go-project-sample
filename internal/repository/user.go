package repository

import "fmt"

const (
	userTable = "users"

	userIdField    = "user_id"
	userLoginField = "login"
)

func (r *Repository) SetNewUser(id int64, login string) (int, error) {
	var innerId int

	query := fmt.Sprintf("INSERT INTO %s (%s, %s) values ($1, $2) RETURNING id",
		userTable, userIdField, userLoginField)
	row := r.db.QueryRow(query, id, login)

	err := row.Scan(&innerId)

	return innerId, err
}

func (r *Repository) GetUseById(id int64) (User, error) {
	var user User

	query := fmt.Sprintf(
		"SELECT %s, %s FROM %s WHERE %s = $1",
		userIdField, userLoginField, userTable, userIdField)

	err := r.db.Get(&user, query, id)

	return user, err
}
