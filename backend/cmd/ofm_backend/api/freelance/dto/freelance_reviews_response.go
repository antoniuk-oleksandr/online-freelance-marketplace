package dto

import "ofm_backend/cmd/ofm_backend/api/freelance/model"

type FreelanceReviewsResponse struct {
	Reviews        *[]model.Review `json:"reviews"`
	HasMoreReviews bool            `json:"hasMoreReviews"`
	ReviewsCursor  *string         `json:"reviewsCursor"`
}
