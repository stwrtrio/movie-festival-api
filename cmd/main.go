package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/routes"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database and Redis
	config.InitDB()
	config.InitRedis()

	// Setup Echo server
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())

	// Initialize routes
	routes.InitRoutes(e)

	// Start server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
