package repository

type User struct {
	Id    int64  `db:"user_id"`
	Login string `db:"login"`
}
