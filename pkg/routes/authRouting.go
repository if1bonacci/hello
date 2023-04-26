package authRouting

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
	"github.com/labstack/echo"
)

type UserType struct {
	Id       string
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

var User = new(UserType)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("users", register)
	e.POST("users", login)
}

func register(ctx echo.Context) (err error) {
	if err = ctx.Bind(User); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	User.Id = uuid.New().String()
	hash, _ := hasher.HashPassword(User.Password)
	User.Password = hash

	r := &RegisterResponse{
		Id:       User.Id,
		UserName: User.UserName,
	}

	return ctx.JSON(http.StatusOK, r)
}

func login(ctx echo.Context) (err error) {
	u := new(UserType)

	if err = ctx.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !hasher.CheckPasswordHash(u.Password, User.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username/password")
	}

	return ctx.JSON(http.StatusOK, LoginResponse{"ws://fancy-chat.io/ws&token=one-time-token"})
}
