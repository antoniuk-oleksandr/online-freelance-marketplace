package routes

import (
	"ofm_backend/cmd/ofm_backend/api/chat/controller"
	"ofm_backend/cmd/ofm_backend/api/chat/repository"
	"ofm_backend/cmd/ofm_backend/api/chat/service"
	websocketmanager "ofm_backend/cmd/ofm_backend/api/chat/websocket_manager"
	"ofm_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterChatRoutes(
	apiGroup fiber.Router, db *sqlx.DB, middleware middleware.Middleware,
) {
	chatRepistory := repository.NewChatRepository(db)
	wsManager := websocketmanager.NewWebSocketManager()
	chatService := service.NewChatService(chatRepistory, wsManager)
	chatController := controller.NewChatController(chatService)

	apiGroup.Get("/ws", middleware.ProcessWebSocketJWT(), chatController.ConnectToWS)
}
