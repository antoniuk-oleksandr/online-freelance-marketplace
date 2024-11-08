package dto

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
	"time"
)

type Freelance struct {
	ID           int64                       `json:"id"`
	CreatedAt    time.Time                   `json:"createdAt"`
	Description  string                      `json:"description"`
	ReviewsCount int64                       `json:"reviewsCount"`
	Rating       float64                     `json:"rating"`
	Title        string                      `json:"title"`
	Images       *[]string                   `json:"images"`
	Category     *model.Category             `json:"category"`
	Packages     *[]Package                  `json:"packages"`
	Freelancer   *FreelanceServiceFreelancer `json:"freelancer"`
	Reviews      *[]model.Review             `json:"reviews"`
}
