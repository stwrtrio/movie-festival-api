package handlers

import (
	"net/http"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/services"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	service services.MovieService
}

func NewMovieHandler(service services.MovieService) *MovieHandler {
	return &MovieHandler{service}
}

func (h *MovieHandler) CreateMovie(c echo.Context) error {
	ctx := c.Request().Context()
	var movieRequest models.MovieRequest
	if err := c.Bind(&movieRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	movie := models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       movieRequest.Title,
		Description: movieRequest.Description,
		Duration:    movieRequest.Duration,
		Genre:       movieRequest.Genre,
		WatchURL:    movieRequest.WatchURL,
		Artist:      movieRequest.Artist,
	}

	err := h.service.CreateMovie(ctx, &movie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Successfully created movie"})
}
