package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"github.com/stwrtrio/movie-festival-api/internal/routes"
	"github.com/stwrtrio/movie-festival-api/internal/schedulers"
	"github.com/stwrtrio/movie-festival-api/internal/services"
	"github.com/stwrtrio/movie-festival-api/pkg/kafka"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database and Redis
	config.InitDB()
	config.InitRedis()

	// Initialize Kafka configuration
	kafkaConfig := config.LoadKafkaConfig()

	// Init producer
	producer, err := kafka.NewProducer(kafkaConfig.Brokers)
	if err != nil {
		log.Fatal("Failed to create Kafka producer:", err)
	}
	defer producer.Close()

	// Setup Echo server
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())

	// Initialize repositories
	movieRepo := repositories.NewMovieRepository(config.DB)
	ratingRepo := repositories.NewRatingRepository(config.DB)

	// Initialize services
	movieService := services.NewMovieService(movieRepo)
	ratingService := services.NewRatingService(ratingRepo, producer, kafkaConfig.Topic)

	// Initialize scheduler
	scheduler := schedulers.NewScheduler()

	scheduler.Start(10 * time.Second)

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieService)
	ratingHandler := handlers.NewRatingHandler(ratingService)

	// Initialize routes
	routes.InitRoutes(e, movieHandler, ratingHandler)

	// Start server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
