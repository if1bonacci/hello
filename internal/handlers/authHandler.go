package handlers

import (
	"net/http"
	"time"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
	"github.com/if1bonacci/lets-go-chat/pkg/tokenGenerator"
	"github.com/labstack/echo"
)

type AuthRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

type LoginResponse struct {
	Url string `json:"url"`
}

const ChatLink = "ws://fancy-chat.io/ws&token="

func Register(ctx echo.Context) (err error) {
	request := new(AuthRequest)

	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := repositories.CreateUser(request.UserName, request.Password)
	repositories.StoreUser(*user)

	resp := &RegisterResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func Login(ctx echo.Context) (err error) {
	u := new(AuthRequest)

	if err = ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := repositories.GetUserByName(u.UserName)

	if !hasher.CheckPasswordHash(u.Password, user.Password) || u.UserName != user.UserName {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username/password")
	}

	ctx.Response().Header().Set("X-Rate-Limit", "3000")
	ctx.Response().Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())

	return ctx.JSON(http.StatusOK, LoginResponse{ChatLink + tokenGenerator.Generate()})
}
