package dto

type FreelanceByIDResponse struct {
	Service        *Freelance `json:"service"`
	HasMoreReviews bool       `json:"hasMoreReviews"`
	ReviewsCursor  *string    `json:"reviewsCursor"`
}
