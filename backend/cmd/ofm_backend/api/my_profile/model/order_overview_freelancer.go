package model

type OrderOverviewFreelancer struct {
	Username string `json:"username" db:"username"`
	Id int `json:"id" db:"id"`
	Avatar string `json:"avatar" db:"avatar"`
}
