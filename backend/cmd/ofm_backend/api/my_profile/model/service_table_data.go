package model

import "time"

type ServiceTableData struct {
	Id          int       `json:"service_id" db:"service_id"`
	Title       string    `json:"title" db:"title"`
	Price       float64   `json:"price" db:"price"`
	Category    string    `json:"category" db:"category"`
	Rating      float64   `json:"rating" db:"rating"`
	OrdersCount int       `json:"orders_count" db:"orders_count"`
	Date        time.Time `json:"date" db:"date"`
	Image       string    `json:"image" db:"image"`
}
