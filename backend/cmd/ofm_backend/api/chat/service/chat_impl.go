package service

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/chat/helpers"
	"ofm_backend/cmd/ofm_backend/api/chat/model"
	"ofm_backend/cmd/ofm_backend/api/chat/repository"
	websocketmanager "ofm_backend/cmd/ofm_backend/api/chat/websocket_manager"
	"ofm_backend/cmd/ofm_backend/enum"

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

		var websocketRequest model.WebSocketRequest
		err = json.Unmarshal(msg, &websocketRequest)
		if err != nil {
			break
		}

		switch websocketRequest.Type {
		case enum.ChatMessage:
			err = chatServ.HandleChatMessage(mt, websocketRequest.Data, userId, conn)
			if err != nil {
				break
			}
		}
	}
}

func (chatServ chatService) HandleChatMessage(
	mt int, messageData []byte,
	userId int, conn *websocket.Conn,
) error {
	var chatMessage model.ChatMessageToReceive
	err := json.Unmarshal(messageData, &chatMessage)
	if err != nil {
		return err
	}

	messageId, err := chatServ.chatRepository.SaveMessage(chatMessage)
	if err != nil {
		return err
	}

	helpers.PrepareChatMessage(&chatMessage, messageId, userId)

	chatMessageBytes, err := json.Marshal(chatMessage)
	if err != nil {
		return err
	}

	websocketRequest := model.WebSocketRequest{
		Type: enum.ChatMessage,
		Data: chatMessageBytes,
	}

	messageBytes, err := json.Marshal(websocketRequest)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(mt, messageBytes); err != nil {
		return err
	}

	//send to partner
	receiver := chatMessage.ChatPartnerId
	chatMessage.ChatPartnerId = userId

	chatMessageBytes, err = json.Marshal(chatMessage)
	if err != nil {
		return err
	}

	websocketRequest = model.WebSocketRequest{
		Type: enum.ChatMessage,
		Data: chatMessageBytes,
	}

	messageBytes, err = json.Marshal(websocketRequest)
	if err != nil {
		return err
	}

	partnerConn := chatServ.wsManager.GetConnection(receiver)
	if partnerConn != nil {
		err = partnerConn.WriteMessage(mt, messageBytes)
		if err != nil {
			return err
		}
	}

	return nil
}
