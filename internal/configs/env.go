package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvInterface interface {
	GetMongoUri() string
	GetDbName() string
}

type Env struct {
	mongoUri string
	dbName   string
}

func ProvideEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Env{
		mongoUri: os.Getenv("MONGOURI"),
		dbName:   os.Getenv("DB_NAME"),
	}
}

func (e *Env) GetMongoUri() string {
	return e.mongoUri
}

func (e *Env) GetDbName() string {
	return e.dbName
}
