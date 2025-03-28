package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/config"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database and Redis
	config.InitDB()
	config.InitRedis()

	// Setup Echo server
	e := echo.New()

	// Start server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
