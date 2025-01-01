package dto

type ReviewsResponse struct {
	Reviews        *[]UserByIdReviewDto `json:"reviews"`
	HasMoreReviews bool                 `json:"hasMoreReviews"`
	ReviewsCursor  *string              `json:"reviewsCursor"`
}
