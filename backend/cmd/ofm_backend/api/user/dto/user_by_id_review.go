package dto

import (
	"time"
)

type UserByIdReviewDto struct {
	ID            int64         `json:"id"`
	Rating        int           `json:"rating"`
	Content       string        `json:"content"`
	CreatedAt     time.Time     `json:"createdAt"`
	EndedAt       time.Time     `json:"endedAt"`
	Customer      ReviewUserDTO `json:"customer"`
	ReviewService ReviewService `json:"service"`
}
