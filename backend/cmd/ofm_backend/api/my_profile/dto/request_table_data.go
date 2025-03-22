package dto

import "time"

type RequestTableData struct {
	Id                int    `json:"id"`
	Title             string    `json:"title"`
	Status            int       `json:"status"`
	Price             float64   `json:"price"`
	CustomerFirstName string    `json:"customerFirstName"`
	CustomerSurname   string    `json:"customerSurname"`
	CustomerAvatar    string    `json:"customerAvatar"`
	Date              time.Time `json:"date"`
}
