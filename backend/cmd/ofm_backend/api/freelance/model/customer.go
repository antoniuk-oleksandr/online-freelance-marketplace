package model

type Customer struct {
	ID        int64            `json:"user_id" db:"user_id"`
	Username  string           `json:"username" db:"username"`
	Avatar    string           `json:"avatar" db:"avatar"`
}
