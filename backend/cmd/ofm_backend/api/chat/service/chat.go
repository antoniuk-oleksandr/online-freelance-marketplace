package service

import (
	"github.com/gofiber/contrib/websocket"
)

type ChatService interface {
	HandleWSConnection(conn *websocket.Conn, userId int)
	HandleChatMessage(mt int, msg []byte, userId int, conn *websocket.Conn) error
}
