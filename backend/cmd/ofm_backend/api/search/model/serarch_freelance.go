package model

import "time"

type SearchService struct {
	ID                            int64     `json:"service_id" db:"service_id"`
	CreatedAt                     time.Time `json:"created_at" db:"created_at"`
	Description                   string    `json:"description" db:"description"`
	Title                         string    `json:"title" db:"title"`
	CategoryId                    int64     `json:"category_id" db:"category_id"`
	FreelancerId                  int64     `json:"freelancer_id" db:"freelancer_id"`
	Image                         *string   `json:"image" db:"image"`
	ReviewsCount                  int64     `json:"reviewsCount" db:"reviews_count"`
	Rating                        float64   `json:"rating" db:"rating"`
	MinPrice                      float64   `json:"minPrice" db:"min_price"`
	LastMonthCompletedOrdersCount int64     `db:"last_month_completed_orders_count"`
	Level                         float64   `db:"level"`
}
