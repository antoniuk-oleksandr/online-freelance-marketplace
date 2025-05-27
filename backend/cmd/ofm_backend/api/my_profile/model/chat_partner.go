package model

import "time"

type ChatPartner struct {
	Id         int64     `json:"user_id" db:"user_id"`
	Username   string    `json:"username" db:"username"`
	Avatar     string    `json:"avatar" db:"avatar"`
	LastSeenAt time.Time `json:"last_seen_at" db:"last_seen_at"`
}
