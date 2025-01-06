package model

type BestReview struct {
	FirstName string `json:"first_name" db:"first_name"`
	Surname   string `json:"surname" db:"surname"`
	Content   string `json:"content" db:"content"`
	Rating    int    `json:"rating" db:"rating"`
}
