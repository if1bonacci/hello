package repositories

import (
	"sync"

	"github.com/if1bonacci/lets-go-chat/internal/models"
)

type chatRepository struct {
	chatUsers map[string]models.User
}

var instance *chatRepository
var once sync.Once

func NewChat() *chatRepository {
	once.Do(func() {
		instance = &chatRepository{
			chatUsers: make(map[string]models.User),
		}
	})

	return instance
}

func (rep *chatRepository) IsActive(token string) bool {
	_, found := rep.chatUsers[token]

	return found
}

func (rep *chatRepository) Add(user models.User) {
	rep.chatUsers[user.Token] = user
}

func (rep *chatRepository) List() map[string]models.User {
	return rep.chatUsers
}
