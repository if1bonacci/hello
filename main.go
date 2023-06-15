package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/if1bonacci/lets-go-chat/docs"
)

// @title           Lets go chat
// @version         1.0
// @description     This is a go sandbox.

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Echo instance
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
