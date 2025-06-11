package controller

import (
	"ofm_backend/cmd/ofm_backend/api/chat/service"
	"ofm_backend/cmd/ofm_backend/utils"

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
	userId, err := utils.GetUserIdFromLocals("userId", ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ErrUserNotFound.Error(),
		})
	}

	return websocket.New(func(conn *websocket.Conn) {
		chatController.chatService.HandleWSConnection(conn, userId)
	})(ctx)
}