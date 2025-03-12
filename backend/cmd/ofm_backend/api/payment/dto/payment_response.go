package dto

type PaymentResponse struct {
	Success bool  `json:"success"`
	OrderId int64 `json:"orderId"`
}
