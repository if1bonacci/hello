package handlers

import (
	"fmt"
	"net/http"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo"
)

type MessageRequest struct {
	Body string `json:"body"`
}

func CreateMessage(ctx echo.Context) (err error) {
	request := new(MessageRequest)

	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	repositories.InitMessage().Add(request.Body)

	return ctx.JSON(http.StatusOK, "success")
}

func AllMessages(ctx echo.Context) (err error) {
	messages := repositories.InitMessage().List()

	for _, message := range messages {
		fmt.Println(message.Body)
	}

	return ctx.JSON(http.StatusOK, messages)
}
