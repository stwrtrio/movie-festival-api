package handlers

import (
	"net/http"

	"github.com/stwrtrio/movie-festival-api/internal/services"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	service services.MovieService
}

func NewMovieHandler(service services.MovieService) *MovieHandler {
	return &MovieHandler{service}
}

func (h *MovieHandler) CreateMovie(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
}
