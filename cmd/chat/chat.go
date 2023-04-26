package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	// routes.InitUserRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
