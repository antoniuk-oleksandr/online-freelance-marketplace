package model

type FreelanceServiceFreelancer struct {
	ID           int64   `json:"id" db:"user_id"`
	Username     string  `json:"username" db:"username"`
	FirstName    string  `json:"first_name" db:"first_name"`
	Surname      string  `json:"surname" db:"surname"`
	Avatar       string  `json:"avatar" db:"avatar"`
	Rating       float64 `json:"rating" db:"rating"`
	Level        float64 `json:"level" db:"level"`
	ReviewsCount int64   `json:"reviews_count" db:"reviews_count"`
}
