package model

import "time"

type RequestTableData struct {
	Id                int    `json:"order_id"`
	Title             string    `json:"title"`
	Status            int       `json:"status"`
	Price             float64   `json:"price"`
	CustomerFirstName string    `json:"customer_first_name"`
	CustomerSurname   string    `json:"customer_surname"`
	CustomerAvatar    string    `json:"customer_avatar"`
	Date              time.Time `json:"date"`
}
