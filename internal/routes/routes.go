package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
)

// InitRoutes to initializes the routes for the application
func InitRoutes(e *echo.Echo) {
	// endpoint health check
	e.GET("/health-check", handlers.HealthCheckHandler)
}
