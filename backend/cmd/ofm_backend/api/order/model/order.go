package model

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type Order struct {
	Id       int64       `json:"order_id" db:"order_id"`
	CratedAt time.Time   `json:"created_at" db:"created_at"`
	Status   enum.Status `json:"status_id" db:"status_id"`
}
