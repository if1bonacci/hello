package handlers

import (
	"fmt"
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

func ActiveUsers(ctx echo.Context) (err error) {
	fmt.Println(repositories.NewChat().List())
	return ctx.JSON(http.StatusOK, repositories.NewChat().List())
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

	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		message := "client " + user.Id + " Connected!"
		log.Println(message)
		reader(conn)
	}
}

func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(p))

	for _, chat := range repositories.NewChat().List() {
		if err := chat.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}