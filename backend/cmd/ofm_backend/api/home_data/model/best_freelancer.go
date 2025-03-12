package model

type BestFreelance struct {
	Id          int64  `json:"id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Image       *string `json:"image" db:"image"`
}
