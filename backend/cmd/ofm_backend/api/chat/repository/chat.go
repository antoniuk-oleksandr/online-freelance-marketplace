package repository

import (
	"ofm_backend/cmd/ofm_backend/api/chat/model"
)

type ChatRepository interface {
	SaveMessage(message model.ChatMessageToReceive) (int, error)
	
}
