package routing

import (
	"github.com/if1bonacci/lets-go-chat/internal/handlers"
	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("user", handlers.Register)
	e.POST("user/login", handlers.Login)
}
