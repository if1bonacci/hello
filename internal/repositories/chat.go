package repositories

import (
	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/models"
)

type ChatRepository struct {
	chatUsers map[string]*websocket.Conn
}

func ProvideChatRepo() ChatRepository {
	return ChatRepository{
		chatUsers: make(map[string]*websocket.Conn),
	}
}

func (rep *ChatRepository) IsActive(token string) bool {
	_, found := rep.chatUsers[token]

	return found
}

func (rep *ChatRepository) Add(user models.User, conn *websocket.Conn) {
	rep.chatUsers[user.Token] = conn
}

func (rep *ChatRepository) Remove(token string) {
	delete(rep.chatUsers, token)
}

func (rep *ChatRepository) List() map[string]*websocket.Conn {
	return rep.chatUsers
}
