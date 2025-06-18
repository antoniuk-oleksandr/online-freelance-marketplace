package model

type OrderDeliveryFreelancer struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
