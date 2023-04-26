package main

import (
	"github.com/if1bonacci/lets-go-chat/pkg/routes/authRouting"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	authRouting.InitAuthRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
