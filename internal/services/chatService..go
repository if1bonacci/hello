package services

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type ChatService struct {
	upgrader websocket.Upgrader
}

func NewChatService() *ChatService {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &ChatService{
		upgrader: upgrader,
	}
}
