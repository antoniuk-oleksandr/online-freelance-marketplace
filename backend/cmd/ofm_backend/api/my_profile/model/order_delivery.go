package model

import "ofm_backend/cmd/ofm_backend/enum"

type OrderDelivery struct {
	Status       enum.Status                `json:"status" db:"status"`
	DeliveryData *OrderDeliveryData         `json:"delivery" db:"delivery"`
	Freelancer   *OrderDeliveryFreelancer   `json:"freelancer" db:"freelancer"`
	Payment      *OrderDeliveryPayment      `json:"payment" db:"payment"`
	Cancellation *OrderDeliveryCancellation `json:"cancellation" db:"cancellation"`
}
