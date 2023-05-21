package repositories

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/models"
)

type chatRepository struct {
	chatUsers map[string]*websocket.Conn
}

var instance *chatRepository
var once sync.Once

func NewChat() *chatRepository {
	once.Do(func() {
		instance = &chatRepository{
			chatUsers: make(map[string]*websocket.Conn),
		}
	})

	return instance
}

func (rep *chatRepository) IsActive(token string) bool {
	_, found := rep.chatUsers[token]

	return found
}

func (rep *chatRepository) Add(user models.User, conn *websocket.Conn) {
	rep.chatUsers[user.Token] = conn
}

func (rep *chatRepository) List() map[string]*websocket.Conn {
	return rep.chatUsers
}
