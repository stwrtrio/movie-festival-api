package handlers

import (
	"net/http"
	"strconv"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/services"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"
	"github.com/stwrtrio/movie-festival-api/pkg/validator"

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

	if err := validator.ValidateStruct(movieRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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

func (h *MovieHandler) UpdateMovie(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "movie ID is required"})
	}

	var movieRequest models.MovieRequest
	if err := c.Bind(&movieRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := validator.ValidateStruct(movieRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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

	err := h.service.UpdateMovie(ctx, &movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update movie"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "movie updated successfully"})
}

func (h *MovieHandler) GetMovies(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	pagination := models.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}

	useCache := false
	if c.QueryParam("use_cache") == "true" || c.QueryParam("use_cache") == "1" {
		useCache = true
	}

	result, err := h.service.GetMovies(c.Request().Context(), pagination, useCache)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
