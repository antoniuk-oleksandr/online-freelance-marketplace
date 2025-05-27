package model

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type ChatMessage struct {
	Id       int64                `json:"chat_message_id" db:"chat_message_id"`
	SenderId int64                `json:"sender_id" db:"sender_id"`
	Content  string               `json:"content" db:"content"`
	SentAt   time.Time            `json:"sent_at" db:"sent_at"`
	Files    []string             `json:"files" db:"files"`
	Type     enum.ChatMessageType `json:"type" db:"type"`
}
