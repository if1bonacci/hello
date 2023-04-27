package main

import (
	"fmt"
	"os"

	"github.com/if1bonacci/lets-go-chat/internal/routing"
	"github.com/joho/godotenv"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	routing.InitAuthRoutes(e)

	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)

	// Start server
	e.Logger.Fatal(e.Start(address))
}
