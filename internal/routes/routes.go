package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
)

// InitRoutes to initializes the routes for the application
func InitRoutes(e *echo.Echo, movieHandler *handlers.MovieHandler) {
	// endpoint health check
	e.GET("/health-check", handlers.HealthCheckHandler)

	// Grouping the routes
	v1 := e.Group("/api/v1")

	// endpoint for user authentication
	InitAuthRoutes(v1)

	// endpoint for movie
	InitMovieRoutes(v1, movieHandler)
}
