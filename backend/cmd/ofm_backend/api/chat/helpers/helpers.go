package helpers

import (
	"ofm_backend/cmd/ofm_backend/api/chat/model"
	"ofm_backend/cmd/ofm_backend/enum"
)

func PrepareChatMessage(
	chatMessage *model.ChatMessageToReceive,
	messageId int,
	userId int,
) {
	chatMessage.MessageId = messageId
	chatMessage.SenderId = userId
	chatMessage.Type = enum.Read
}
