package app

type UserAPIRequest struct {
	Login string `json:"login"`
	Id    int64  `json:"id"`
}
