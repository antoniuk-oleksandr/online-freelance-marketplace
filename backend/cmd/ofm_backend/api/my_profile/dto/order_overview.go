package dto

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type OrderOverview struct {
	Id                      int                     `json:"id" db:"id"`
	DeliveryDate            *time.Time              `json:"deliveryDate"`
	CreatedAt               time.Time               `json:"createdAt"`
	Subtotal                float64                 `json:"subtotal"`
	ServiceFee              float64                 `json:"serviceFee"`
	TotalPrice              float64                 `json:"totalPrice"`
	Status                  enum.Status             `json:"status"`
	OrderOverviewService    OrderOverviewService    `json:"service"`
	OrderOverviewFreelancer OrderOverviewFreelancer `json:"freelancer"`
}
