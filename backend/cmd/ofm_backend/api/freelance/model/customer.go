package model

type Customer struct {
	ID        int64            `json:"id" db:"user_id"`
	Username  string           `json:"username" db:"username"`
	Avatar    string           `json:"avatar" db:"avatar"`
}
