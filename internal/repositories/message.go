package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/if1bonacci/lets-go-chat/internal/configs"
	"github.com/if1bonacci/lets-go-chat/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageRepositoryInterface interface {
	InitMessage() *mongo.Collection
	Add(mes string)
	Create(mes string) models.Message
	List() []models.Message
}

type MessageRepository struct {
	db configs.MongoDB
}

func ProvideMessageRepo(db configs.MongoDB) MessageRepository {
	return MessageRepository{
		db: db,
	}
}

func (rep *MessageRepository) InitMessage() *mongo.Collection {
	return rep.db.GetCollection("messages")
}

func (rep *MessageRepository) Add(mes string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := models.Message{
		Id:   primitive.NewObjectID(),
		Body: mes,
	}

	rep.InitMessage().InsertOne(ctx, message)
}

func (rep *MessageRepository) Create(mes string) models.Message {
	return models.Message{
		Id:   primitive.NewObjectID(),
		Body: mes,
	}
}

func (rep *MessageRepository) List() []models.Message {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var messages []models.Message
	defer cancel()

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"_id", -1}})
	results, err := rep.InitMessage().Find(context.TODO(), filter, opts)

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
