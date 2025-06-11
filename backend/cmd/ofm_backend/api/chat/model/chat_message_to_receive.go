package model

import "ofm_backend/cmd/ofm_backend/enum"

type ChatMessageToReceive struct {
	MessageId     int                  `json:"messageId"`
	Type          enum.ChatMessageType `json:"type"`
	Content       []byte               `json:"content"`
	ContentIV     []byte               `json:"contentIV"`
	ChatPartnerId int                  `json:"chatPartnerId"`
	SentAt        int64                `json:"sentAt"`
	SenderId      int                  `json:"senderId"`
	OrderId       int                  `json:"orderId"`
}
