package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	//run database
	db, err := InitializeDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.ConnectDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes, err := InitializeRouting()
	if err != nil {
		e.Logger.Fatal(err)
	}
	routes.InitAuthRoutes(e)

	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)

	// Start server
	e.Logger.Fatal(e.Start(address))
}
