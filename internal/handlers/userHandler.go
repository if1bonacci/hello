package handlers

import (
	"net/http"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo"
)

func ListOfUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, repositories.ListOfUsers())
}

func ActiveUsers(ctx echo.Context) (err error) {
	return ctx.JSON(http.StatusOK, repositories.NewChat().List())
}
