package service

import (
	"log"
	"ofm_backend/cmd/ofm_backend/api/chat/repository"
	websocketmanager "ofm_backend/cmd/ofm_backend/api/chat/websocket_manager"

	"github.com/gofiber/contrib/websocket"
)

type chatService struct {
	chatRepository repository.ChatRepository
	wsManager      websocketmanager.WebSocketManager
}

func NewChatService(
	chatRepository repository.ChatRepository,
	wsManager websocketmanager.WebSocketManager,
) ChatService {
	return &chatService{
		chatRepository: chatRepository,
		wsManager:      wsManager,
	}
}

func (chatServ *chatService) HandleWSConnection(conn *websocket.Conn, userId int) {
	chatServ.wsManager.AddConnection(userId, conn)
	defer chatServ.wsManager.RemoveConnection(userId)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		log.Printf("recv: %s", msg)

		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
