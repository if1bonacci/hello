package repositories

import (
	"github.com/google/uuid"
	"github.com/if1bonacci/lets-go-chat/internal/models"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
)

var users = make(map[string]*models.User)

func StoreUser(user *models.User) {
	users[user.UserName] = user
}

func GetUserByName(userName string) *models.User {
	return users[userName]
}

func GetUserByToken(token string) *models.User {
	for _, user := range users {
		if user.Token == token {
			return user
		}
	}

	return nil
}

func RemoveToken(user *models.User) {
	users[user.UserName].Token = ""
}

func ListOfUsers() map[string]*models.User {
	return users
}

func CreateUser(userName string, password string) *models.User {
	hash, _ := hasher.HashPassword(password)

	return &models.User{
		Id:       uuid.New().String(),
		UserName: userName,
		Password: hash,
		Token:    "",
	}
}
