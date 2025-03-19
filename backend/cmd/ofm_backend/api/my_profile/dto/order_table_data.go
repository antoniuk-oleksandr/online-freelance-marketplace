package dto

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type OrderTableData struct {
	Id     int         `json:"id"`
	Title  string      `json:"title"`
	Status enum.Status `json:"status"`
	Price  float64     `json:"price"`
	Date   time.Time   `json:"date"`
	Image  string      `json:"image"`
}
