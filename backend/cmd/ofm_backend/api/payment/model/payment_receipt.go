package model

import "time"

type PaymentReceipt struct {
	OrderId       int64     `json:"order_id" db:"order_id"`
	ServiceName   string    `json:"service_name" db:"service_name"`
	PackageName   string    `json:"package_name" db:"package_name"`
	Price         float64   `json:"price" db:"price"`
	ServiceFee    float64   `json:"service_fee" db:"service_fee"`
	Total         float64   `json:"total" db:"total"`
	Date          time.Time `json:"date" db:"date"`
	CustomerEmail string    `json:"customer_email" db:"customer_email"`
}
