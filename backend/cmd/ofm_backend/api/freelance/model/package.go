package model

type Package struct {
	ID           int64   `json:"id" db:"id"`
	DeliveryDays int64   `json:"delivery_days" db:"delivery_days"`
	Description  string  `json:"description" db:"description"`
	Price        float64 `json:"price" db:"price"`
	Title        string  `json:"title" db:"title"`
}
