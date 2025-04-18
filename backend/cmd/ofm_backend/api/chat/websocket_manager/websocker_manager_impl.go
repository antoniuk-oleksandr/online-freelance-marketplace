package websocketmanager

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type webSocketManager struct {
	mu          sync.RWMutex
	connections map[int]*websocket.Conn
}

func NewWebSocketManager() WebSocketManager {
	return &webSocketManager{
		connections: make(map[int]*websocket.Conn),
	}
}

func (w *webSocketManager) AddConnection(id int, conn *websocket.Conn) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.connections[id] = conn
}

func (w *webSocketManager) GetConnection(id int) *websocket.Conn {
	return w.connections[id]
}

func (w *webSocketManager) RemoveConnection(id int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.connections, id)
}
