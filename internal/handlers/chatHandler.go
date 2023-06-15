package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	repo     repositories.MessageRepository
	userRepo *repositories.UserRepository
	chatRepo *repositories.ChatRepository
}

func ProvideChatHandler(
	r repositories.MessageRepository,
	u *repositories.UserRepository,
	c *repositories.ChatRepository,
) ChatHandler {
	return ChatHandler{
		repo:     r,
		userRepo: u,
		chatRepo: c,
	}
}

func (h *ChatHandler) WebSocket(ctx echo.Context) (err error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	token := ctx.QueryParam("token")
	user := h.userRepo.GetUserByToken(token)
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User by token is not defined.")
	}

	if h.chatRepo.IsActive(token) {
		return echo.NewHTTPError(http.StatusBadRequest, "User is already active.")
	}

	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	h.chatRepo.Add(*user, conn)
	messages := h.repo.List()
	for _, mes := range messages {
		err = conn.WriteMessage(1, []byte(mes.Body))
		if err != nil {
			log.Println(err)
		}
	}
	h.userRepo.RemoveToken(user)

	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		reader(conn, token, h.repo, *h.chatRepo)
	}
}

func reader(
	conn *websocket.Conn,
	token string,
	repo repositories.MessageRepository,
	chatRepo repositories.ChatRepository,
) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)

		chatRepo.Remove(token)
		return
	}

	go repo.Add(string(p))

	for _, chat := range chatRepo.List() {
		if err := chat.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
