package service

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/chat/dto"
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

		var chatMessage dto.ChatMessageToReceive
		err = json.Unmarshal(msg, &chatMessage)
		if err != nil {
			break
		}

		partnerConn := chatServ.wsManager.GetConnection(chatMessage.ChatPartnerId)
		if partnerConn != nil {
			err = partnerConn.WriteMessage(mt, msg)
			if err != nil {
				break
			}
		}
		
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			break
		}
	}
}
