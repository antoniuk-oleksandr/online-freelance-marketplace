package websocketmanager

import "github.com/gofiber/contrib/websocket"

type WebSocketManager interface {
	AddConnection(id int, conn *websocket.Conn)
	RemoveConnection(id int)
	GetConnection(id int) *websocket.Conn
}
