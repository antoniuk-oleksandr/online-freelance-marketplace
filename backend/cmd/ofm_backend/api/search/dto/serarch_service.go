package dto

import "time"

type SearchFreelance struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	Title        string    `json:"title"`
	CategoryId   int64     `json:"category_id"`
	FreelancerId int64     `json:"freelancer_id"`
	Image        *string   `json:"image"`
	ReviewsCount int64     `json:"reviewsCount"`
	Rating       float64   `json:"rating"`
	MinPrice     float64   `json:"minPrice"`
}
