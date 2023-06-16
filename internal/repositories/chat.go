package repositories

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/models"
)

type ChatRepositoryInterface interface {
	IsActive(token string) bool
	Add(user models.User, conn *websocket.Conn)
	Remove(token string)
	List() map[string]*websocket.Conn
}

type ChatRepository struct {
	chatUsers map[string]*websocket.Conn
}

var chatRepoInst *ChatRepository
var chatRepoOnce sync.Once

func ProvideChatRepo() *ChatRepository {
	chatRepoOnce.Do(func() {
		chatRepoInst = &ChatRepository{
			chatUsers: make(map[string]*websocket.Conn),
		}
	})

	return chatRepoInst
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
