package dto

import (
	"ofm_backend/cmd/ofm_backend/api/user/model"
	"time"
)

type UserByIdTO struct {
	ID           int64                `json:"id"`
	About        *string              `json:"about"`
	CreatedAt    time.Time            `json:"createdAt"`
	FirstName    string               `json:"firstName"`
	Level        float64              `json:"level"`
	Surname      string               `json:"surname"`
	Username     string               `json:"username"`
	Avatar       *string              `json:"avatar"`
	Rating       *float64             `json:"rating"`
	ReviewsCount *int64               `json:"reviewsCount"`
	Skills       *[]model.Skill       `json:"skills"`
	Languages    *[]model.Language    `json:"languages"`
	Reviews      []UserByIdReviewDto  `json:"reviews"`
	Services     []ServiceByIdDto `json:"services"`
}
