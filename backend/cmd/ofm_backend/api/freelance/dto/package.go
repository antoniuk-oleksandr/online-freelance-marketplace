package dto

type Package struct {
	ID           int64   `json:"id"`
	DeliveryDays int64   `json:"deliveryDays"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Title        string  `json:"title"`
}