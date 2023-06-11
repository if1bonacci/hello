package repositories

import (
	"sync"

	"github.com/google/uuid"
	"github.com/if1bonacci/lets-go-chat/internal/models"
)

type messageRepository struct {
	messages map[string]*models.Message
}

var instanceM *messageRepository
var onceM sync.Once

func InitMessage() *messageRepository {
	onceM.Do(func() {
		instanceM = &messageRepository{
			messages: make(map[string]*models.Message),
		}
	})

	return instanceM
}

func (rep *messageRepository) Add(mes *models.Message) {
	rep.messages[mes.Id] = mes
}

func (rep *messageRepository) Create(mes string) *models.Message {
	return &models.Message{
		Id:   uuid.New().String(),
		Body: mes,
	}
}

func (rep *messageRepository) List() map[string]*models.Message {
	return rep.messages
}
