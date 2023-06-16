package handlers

import (
	"fmt"
	"net/http"

	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo/v4"
)

type MessageRequest struct {
	Body string `json:"body"`
}

type MessageHandler struct {
	repo repositories.MessageRepository
}

func ProvideMessageHandler(r repositories.MessageRepository) MessageHandler {
	return MessageHandler{
		repo: r,
	}
}

// CreateMessage godoc
// @Summary      Create new message
// @Description  Create new message
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param input body MessageRequest true "some text"
// @Success      200  {string}  string    "success"
// @Router       /message [post]
func (h *MessageHandler) CreateMessage(ctx echo.Context) (err error) {
	request := new(MessageRequest)

	if err = ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.repo.Add(request.Body)

	return ctx.JSON(http.StatusOK, "success")
}

// AllMessages godoc
// @Summary      List of messages
// @Description  get list of message
// @Tags         messages
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Message
// @Router       /messages [get]
func (h *MessageHandler) AllMessages(ctx echo.Context) (err error) {
	messages := h.repo.List()

	for _, message := range messages {
		fmt.Println(message.Body)
	}

	return ctx.JSON(http.StatusOK, messages)
}
