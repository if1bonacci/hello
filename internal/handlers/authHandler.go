package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/if1bonacci/lets-go-chat/internal/models"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
	"github.com/if1bonacci/lets-go-chat/pkg/tokenGenerator"
	"github.com/labstack/echo"
)

type RegisterResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

type LoginResponse struct {
	Url string `json:"url"`
}

const ChatLink = "ws://fancy-chat.io/ws&token="

func Register(ctx echo.Context) (err error) {
	user := models.User
	if err = ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Id = uuid.New().String()
	hash, _ := hasher.HashPassword(user.Password)
	user.Password = hash

	r := &RegisterResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}

	return ctx.JSON(http.StatusOK, r)
}

func Login(ctx echo.Context) (err error) {
	user := models.User
	u := new(models.UserType)

	if err = ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !hasher.CheckPasswordHash(u.Password, user.Password) || u.UserName != user.UserName {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username/password")
	}

	ctx.Response().Header().Set("X-Rate-Limit", "3000")
	ctx.Response().Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())

	return ctx.JSON(http.StatusOK, LoginResponse{ChatLink + tokenGenerator.Generate()})
}
