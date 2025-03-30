package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/mocks"
	"github.com/stwrtrio/movie-festival-api/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.Movie{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Rating:      8.8,
	}

	jsonMovie, _ := json.Marshal(movie)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies", bytes.NewReader(jsonMovie))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := movieHandler.CreateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateMovieHandler_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.Movie{
		Title:       "",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Rating:      8.8,
	}

	jsonMovie, _ := json.Marshal(movie)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies", bytes.NewReader(jsonMovie))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(errors.New("invalid title"))

	err := movieHandler.CreateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateMovieHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.MovieRequest{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Duration:    148,
		Genre:       "Sci-Fi",
		WatchURL:    "http://example.com/watch",
		Artist:      "Christopher Nolan",
	}

	jsonMovie, _ := json.Marshal(movie)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/movies/123", bytes.NewReader(jsonMovie))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/movies/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	mockService.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := movieHandler.UpdateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateMovieHandler_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.MovieRequest{
		Title:       "",
		Description: "A mind-bending thriller",
		Duration:    148,
		Genre:       "Sci-Fi",
		WatchURL:    "http://example.com/watch",
		Artist:      "Christopher Nolan",
	}

	jsonMovie, _ := json.Marshal(movie)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/movies/123", bytes.NewReader(jsonMovie))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/movies/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	mockService.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := movieHandler.UpdateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
