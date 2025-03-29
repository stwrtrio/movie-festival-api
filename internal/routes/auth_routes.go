package routes

import (
	"github.com/stwrtrio/movie-festival-api/internal/handlers"

	"github.com/labstack/echo/v4"
)

// InitAuthRoutes mengatur routing untuk autentikasi
func InitAuthRoutes(e *echo.Group) {
	authGroup := e.Group("/auth")
	authGroup.POST("/register", handlers.RegisterHandler)
	authGroup.POST("/login", handlers.LoginHandler)
}
