package dto

import "time"

type ServiceTableData struct {
	Id          int       `json:"id"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Rating      float64   `json:"rating"`
	OrdersCount int       `json:"ordersCount"`
	Date        time.Time `json:"date"`
}
