package routes

import (
	"github.com/stwrtrio/movie-festival-api/internal/handlers"

	"github.com/labstack/echo/v4"
)

func InitRatingRoutes(e *echo.Group, ratingHandler *handlers.RatingHandler) {
	e.POST("/movies/rate", ratingHandler.RateMovie)
}
