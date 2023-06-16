package repositories

import (
	"sync"

	"github.com/google/uuid"
	"github.com/if1bonacci/lets-go-chat/internal/models"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
)

type UserRepositoryInterface interface {
	StoreUser(user *models.User)
	GetUserByName(userName string) *models.User
	GetUserByToken(token string) *models.User
	RemoveToken(user *models.User)
	ListOfUsers() map[string]*models.User
	CreateUser(userName string, password string) *models.User
}

type UserRepository struct {
	users map[string]*models.User
}

var instanceU *UserRepository
var onceU sync.Once

func ProvideUserRepo() *UserRepository {
	onceU.Do(func() {
		instanceU = &UserRepository{
			users: make(map[string]*models.User),
		}
	})

	return instanceU
}

func (u *UserRepository) StoreUser(user *models.User) {
	u.users[user.UserName] = user
}

func (u *UserRepository) GetUserByName(userName string) *models.User {
	return u.users[userName]
}

func (u *UserRepository) GetUserByToken(token string) *models.User {
	for _, user := range u.users {
		if user.Token == token {
			return user
		}
	}

	return nil
}

func (u *UserRepository) RemoveToken(user *models.User) {
	u.users[user.UserName].Token = ""
}

func (u *UserRepository) ListOfUsers() map[string]*models.User {
	return u.users
}

func (u *UserRepository) CreateUser(userName string, password string) *models.User {
	hash, _ := hasher.HashPassword(password)

	return &models.User{
		Id:       uuid.New().String(),
		UserName: userName,
		Password: hash,
		Token:    "",
	}
}
