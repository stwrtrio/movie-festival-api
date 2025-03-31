package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"github.com/stwrtrio/movie-festival-api/internal/routes"
	"github.com/stwrtrio/movie-festival-api/internal/services"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database and Redis
	config.InitDB()
	config.InitRedis()

	// Initialize Kafka
	kafkaClient, err := config.InitKafka()
	if err != nil {
		log.Fatalf("Failed to initialize Kafka: %v", err)
	}
	log.Println("Connected to Kafka")
	defer kafkaClient.Close()

	// Setup Echo server
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())

	// Initialize repositories
	movieRepo := repositories.NewMovieRepository(config.DB)

	// Initialize services
	movieService := services.NewMovieService(movieRepo)

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieService)

	// Initialize routes
	routes.InitRoutes(e, movieHandler)

	// Start server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
