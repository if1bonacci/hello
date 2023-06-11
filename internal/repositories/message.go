package repositories

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/if1bonacci/lets-go-chat/internal/configs"
	"github.com/if1bonacci/lets-go-chat/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type messageRepository struct {
	messages *mongo.Collection
}

var instanceM *messageRepository
var onceM sync.Once

func InitMessage() *messageRepository {
	onceM.Do(func() {
		instanceM = &messageRepository{
			messages: configs.GetCollection(configs.DB, "messages"),
		}
	})

	return instanceM
}

func (rep *messageRepository) Add(mes string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := models.Message{
		Id:   primitive.NewObjectID(),
		Body: mes,
	}

	rep.messages.InsertOne(ctx, message)
}

func (rep *messageRepository) Create(mes string) models.Message {
	return models.Message{
		Id:   primitive.NewObjectID(),
		Body: mes,
	}
}

func (rep *messageRepository) List() []models.Message {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var messages []models.Message
	defer cancel()

	results, err := rep.messages.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleMessage models.Message
		if err = results.Decode(&singleMessage); err != nil {
			fmt.Println(err)
		}

		messages = append(messages, singleMessage)
	}

	return messages
}
