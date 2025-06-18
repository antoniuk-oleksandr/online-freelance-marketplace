package model

type OrderReview struct {
	Content     string              `json:"content" db:"content"`
	Rating      int                 `json:"rating" db:"rating"`
	Customer    OrderReviewCustomer `json:"customer" db:"customer"`
	OrderStatus int                 `json:"orderStatus" db:"order_status"`
}
