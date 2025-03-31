package handlers

import (
	"net/http"
	"time"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/services"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"
	"github.com/stwrtrio/movie-festival-api/pkg/validator"

	"github.com/labstack/echo/v4"
)

type RatingHandler struct {
	service services.RatingService
}

func NewRatingHandler(service services.RatingService) *RatingHandler {
	return &RatingHandler{service}
}

func (h *RatingHandler) RateMovie(c echo.Context) error {
	ctx := c.Request().Context()
	var ratingRequest models.RatingRequest
	if err := c.Bind(&ratingRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := validator.ValidateStruct(ratingRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rating := models.Rating{
		ID:        utils.GenerateUUID(),
		UserID:    ratingRequest.UserID,
		MovieID:   ratingRequest.MovieID,
		Score:     ratingRequest.Score,
		Comment:   ratingRequest.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := h.service.RateMovie(ctx, &rating)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully rated movie"})
}
