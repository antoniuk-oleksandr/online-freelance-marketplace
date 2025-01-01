package dto

type UserByIDResponse struct {
	User            *UserByIdTO `json:"user"`
	HasMoreReviews  bool        `json:"hasMoreReviews"`
	ReviewsCursor   *string     `json:"reviewsCursor"`
	HasMoreServices bool        `json:"hasMoreServices"`
	ServicesCursor  *string     `json:"servicesCursor"`
}
