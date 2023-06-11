package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(ctx echo.Context) (err error) {
	token := ctx.QueryParam("token")
	user := repositories.GetUserByToken(token)
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User by token is not defined.")
	}

	if repositories.NewChat().IsActive(token) {
		return echo.NewHTTPError(http.StatusBadRequest, "User is already active.")
	}

	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	repositories.NewChat().Add(*user, conn)
	messages := repositories.InitMessage().List()
	for _, mes := range messages {
		err = conn.WriteMessage(1, []byte(mes.Body))
		if err != nil {
			log.Println(err)
		}
	}
	repositories.RemoveToken(user)

	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		reader(conn, token)
	}
}

func reader(conn *websocket.Conn, token string) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)

		repositories.NewChat().Remove(token)
		return
	}

	go repositories.InitMessage().Add(string(p))

	for _, chat := range repositories.NewChat().List() {
		if err := chat.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
