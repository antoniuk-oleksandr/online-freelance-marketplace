package dto

import (
	"ofm_backend/cmd/ofm_backend/enum"
	"time"
)

type ChatMessage struct {
	Id        int64                `json:"id"`
	SenderId  int64                `json:"senderId"`
	Content   []byte               `json:"content"`
	ContentIV []byte               `json:"contentIV"`
	SentAt    time.Time            `json:"sentAt"`
	Files     []string             `json:"files"`
	Type      enum.ChatMessageType `json:"type"`
}
