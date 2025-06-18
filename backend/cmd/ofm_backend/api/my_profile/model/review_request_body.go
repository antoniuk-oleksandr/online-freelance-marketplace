package model

type ReviewRequestBody struct {
	ReviewId      int    `json:"reviewId"`
	Rating        int    `json:"rating"`
	ReviewMessage string `json:"reviewMessage"`
	UserId        int    `json:"userId"`
	OrderId       int    `json:"orderId"`
}
