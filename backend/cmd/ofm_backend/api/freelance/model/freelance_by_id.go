package model

import (
	"time"
)

type FreelanceByID struct {
	ID           int64                       `json:"id" db:"id"`
	CreatedAt    time.Time                   `json:"created_at" db:"created_at"`
	Description  string                      `json:"description" db:"description"`
	ReviewsCount int64                       `json:"reviews_count" db:"reviews_count"`
	Rating       float64                     `json:"rating" db:"rating"`
	Title        string                      `json:"title" db:"title"`
	Images       *[]string                   `json:"images" db:"images"`
	Category     *Category                   `json:"category" db:"category"`
	Packages     *[]Package                  `json:"packages" db:"packages"`
	Freelancer   *FreelanceServiceFreelancer `json:"freelancer" db:"freelancer"`
}
