package handlers

import (
	"net/http"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo     *repositories.UserRepository
	chatRepo repositories.ChatRepository
}

func ProvideUserHandler(r *repositories.UserRepository) UserHandler {
	return UserHandler{
		repo: r,
	}
}

// ListOfUsers godoc
// @Summary      List of users
// @Description  get list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /user/list [get]
func (u *UserHandler) ListOfUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, u.repo.ListOfUsers())
}

func (u *UserHandler) ActiveUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, u.chatRepo.List())
}
