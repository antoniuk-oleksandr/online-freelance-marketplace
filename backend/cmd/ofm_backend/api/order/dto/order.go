package dto

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type Order struct {
	Id       int64       `json:"id"`
	CratedAt time.Time   `json:"createdAt"`
	Status    enum.Status `json:"status"`
}
