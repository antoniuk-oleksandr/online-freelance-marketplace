package service

import "github.com/gofiber/contrib/websocket"

type ChatService interface {
	HandleWSConnection(conn *websocket.Conn, userId int)
}
