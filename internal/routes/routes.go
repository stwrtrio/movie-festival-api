package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/middlewares"
)

// InitRoutes to initializes the routes for the application
func InitRoutes(e *echo.Echo,
	movieHandler *handlers.MovieHandler,
	ratingHandler *handlers.RatingHandler,
) {
	// endpoint health check
	e.GET("/health-check", handlers.HealthCheckHandler)

	// Grouping the routes
	v1 := e.Group("/api/v1")

	// endpoint for user authentication
	InitAuthRoutes(v1)

	v1.Use(middlewares.AuthMiddleware)

	// endpoint for movie
	InitMovieRoutes(v1, movieHandler)

	// endpoint for rating
	InitRatingRoutes(v1, ratingHandler)
}
