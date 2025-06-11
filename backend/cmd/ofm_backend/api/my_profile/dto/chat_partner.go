package dto

import "time"

type ChatPartner struct {
	Id         int64     `json:"partnerId"`
	Username   string    `json:"username"`
	Avatar     string    `json:"avatar"`
	LastSeenAt time.Time `json:"lastSeenAt"`
}
