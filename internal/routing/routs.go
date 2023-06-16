package routing

import (
	"github.com/if1bonacci/lets-go-chat/internal/handlers"
	"github.com/labstack/echo/v4"
)

type RoutingInterface interface {
	InitAuthRoutes(e *echo.Echo)
}

type Routing struct {
	auth    handlers.AuthHandler
	message handlers.MessageHandler
	chat    handlers.ChatHandler
	user    handlers.UserHandler
}

func ProvideRouting(
	a handlers.AuthHandler,
	m handlers.MessageHandler,
	c handlers.ChatHandler,
	u handlers.UserHandler,
) Routing {
	return Routing{
		auth:    a,
		message: m,
		chat:    c,
		user:    u,
	}
}

func (r *Routing) InitAuthRoutes(e *echo.Echo) {
	e.POST("user", r.auth.Register)
	e.POST("user/login", r.auth.Login)
	e.GET("websoket", r.chat.WebSocket)
	e.GET("user/list", r.user.ListOfUsers)
	e.GET("user/active-users", r.user.ActiveUsers)
	e.POST("message", r.message.CreateMessage)
	e.GET("messages", r.message.AllMessages)
}
