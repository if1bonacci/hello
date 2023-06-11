package routing

import (
	"github.com/if1bonacci/lets-go-chat/internal/handlers"
	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("user", handlers.Register)
	e.POST("user/login", handlers.Login)
	e.GET("user/list", handlers.ListOfUsers)
	e.GET("/websoket", handlers.WebSocket)
	e.GET("/websoket/active-users", handlers.ActiveUsers)
	e.POST("message", handlers.CreateMessage)
	e.GET("messages", handlers.AllMessages)
}
