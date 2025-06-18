package model

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type OrderOverview struct {
	Id                      int                     `json:"id" db:"id"`
	DeliveryDate            *time.Time              `json:"delivery_date" db:"delivery_date"`
	CreatedAt               time.Time               `json:"created_at" db:"created_at"`
	Subtotal                float64                 `json:"subtotal" db:"subtotal"`
	ServiceFee              float64                 `json:"service_fee" db:"service_fee"`
	TotalPrice              float64                 `json:"total_price" db:"total_price"`
	Status                  enum.Status             `json:"status" db:"status"`
	OrderOverviewService    OrderOverviewService    `json:"service"`
	OrderOverviewFreelancer OrderOverviewFreelancer `json:"freelancer"`
}
