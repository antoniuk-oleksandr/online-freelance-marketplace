package model

type UserSignInData struct {
	Id       int    `json:"user_id" db:"user_id"`
	Avatar   string `json:"avatar" db:"avatar"`
	Username string `json:"username" db:"username"`
}
