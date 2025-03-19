package model

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type OrderTableData struct {
	Id     int         `json:"order_id" db:"order_id"`
	Title  string      `json:"title" db:"title"`
	Status enum.Status `json:"status" db:"status"`
	Price  float64     `json:"price" db:"price"`
	Date   time.Time   `json:"date" db:"date"`
	Image  string      `json:"image" db:"image"`
}
