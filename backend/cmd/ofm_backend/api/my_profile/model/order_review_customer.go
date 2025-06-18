package model

type OrderReviewCustomer struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	Surname   string `json:"surname" db:"surname"`
	Avatar    string `json:"avatar" db:"avatar"`
}
