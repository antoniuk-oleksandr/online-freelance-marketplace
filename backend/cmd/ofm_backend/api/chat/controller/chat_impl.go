package controller

import (
	"ofm_backend/cmd/ofm_backend/api/chat/service"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type chatController struct {
	chatService service.ChatService
}

func NewChatController(chatService service.ChatService) ChatController {
	return &chatController{
		chatService: chatService,
	}
}

func (chatController *chatController) ConnectToWS(ctx *fiber.Ctx) error {
	userIdInterface := ctx.Locals("userId")
	if userIdInterface == nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	userId := int(userIdInterface.(float64))

	return websocket.New(func(conn *websocket.Conn) {
		chatController.chatService.HandleWSConnection(conn, userId)
	})(ctx)
}
