package main

import (
	"github.com/if1bonacci/lets-go-chat/internal/routing"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	routing.InitAuthRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
