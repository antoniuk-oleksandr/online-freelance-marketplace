package routes

import (
	"ofm_backend/cmd/ofm_backend/api/chat/controller"
	"ofm_backend/cmd/ofm_backend/api/chat/repository"
	"ofm_backend/cmd/ofm_backend/api/chat/service"
	websocketmanager "ofm_backend/cmd/ofm_backend/api/chat/websocket_manager"
	"ofm_backend/internal/middleware"
	middleware_repo "ofm_backend/internal/middleware/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterChatRoutes(apiGroup fiber.Router, db *sqlx.DB) {
	middlewareRepository := middleware_repo.NewMiddlewareRepository(db)
	middleware := middleware.NewMiddleware(middlewareRepository)

	chatRepistory := repository.NewChatRepository(db)
	wsManager := websocketmanager.NewWebSocketManager()
	chatService := service.NewChatService(chatRepistory, wsManager)
	chatController := controller.NewChatController(chatService)

	apiGroup.Get("/ws", middleware.ProcessRegularJWT(), chatController.ConnectToWS)
}
