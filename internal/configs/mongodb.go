package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInterface interface {
	ConnectDB() *mongo.Client
	GetCollection(collectionName string) *mongo.Collection
}

type MongoDB struct {
	env Env
}

func ProvideMongoDB(e Env) MongoDB {
	return MongoDB{
		env: e,
	}
}

func (m *MongoDB) ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.env.GetMongoUri()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	collection := m.ConnectDB().Database(m.env.GetDbName()).Collection(collectionName)
	return collection
}
