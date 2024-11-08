package dto

import (
	"time"
)

type UserByIdServiceDto struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	ReviewsCount *int64      `json:"reviews_count"`
	Rating       *float64    `json:"rating"`
	MinPrice     float64     `json:"min_price"`
}
