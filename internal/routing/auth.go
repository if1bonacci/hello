package routing

import (
	"github.com/if1bonacci/lets-go-chat/internal/handlers"
	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("user", handlers.Register)
	e.POST("user/login", handlers.Login)
	e.GET("/websoket", handlers.WebSocket)
	e.GET("/websoket/active-users", handlers.ActiveUsers)
}
