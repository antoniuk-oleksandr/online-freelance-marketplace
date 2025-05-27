package dto

import "ofm_backend/cmd/ofm_backend/enum"

type ChatMessageToReceive struct {
	Type          enum.ChatMessageType `json:"type"`
	Content       string               `json:"content"`
	ChatPartnerId int                  `json:"chatPartnerId"`
	SentAt        int64                `json:"sentAt"`
}
