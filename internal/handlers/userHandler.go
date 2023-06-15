package handlers

import (
	"net/http"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo"
)

type UserHandler struct {
	repo     repositories.UserRepository
	chatRepo repositories.ChatRepository
}

func ProvideUserHandler(r repositories.UserRepository) UserHandler {
	return UserHandler{
		repo: r,
	}
}

func (u *UserHandler) ListOfUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, u.repo.ListOfUsers())
}

func (u *UserHandler) ActiveUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, u.chatRepo.List())
}
