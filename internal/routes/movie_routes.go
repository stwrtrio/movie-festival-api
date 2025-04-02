package routes

import (
	"github.com/stwrtrio/movie-festival-api/internal/handlers"

	"github.com/labstack/echo/v4"
)

func InitMovieRoutes(e *echo.Group, movieHandler *handlers.MovieHandler) {
	e.GET("/movies", movieHandler.GetMovies)
	e.POST("/movies", movieHandler.CreateMovie)
	e.PUT("/movies/:id", movieHandler.UpdateMovie)
}
