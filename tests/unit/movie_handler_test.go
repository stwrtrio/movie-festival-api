package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	movie := &models.MovieRequest{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Duration:    120,
		WatchURL:    "http://example.com/watch",
		Artist:      "Christopher Nolan",
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
	movie := &models.MovieRequest{
		Title:       "",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Duration:    120,
		WatchURL:    "http://example.com/watch",
		Artist:      "Christopher Nolan",
	}

	jsonMovie, _ := json.Marshal(movie)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies", bytes.NewReader(jsonMovie))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

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

	err := movieHandler.UpdateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMoviesHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.PaginationResponse{
		Page:       1,
		PageSize:   1,
		TotalItems: 1,
		TotalPages: 1,
		Data: models.Movie{
			Title:       "Inception",
			Description: "A mind-bending thriller",
			Genre:       "Sci-Fi",
			Duration:    120,
			WatchURL:    "http://example.com/watch",
			Artist:      "Christopher Nolan",
		},
	}

	page := 1
	limit := 5
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().GetMovies(gomock.Any(), gomock.Any()).Return(nil)

	err := movieHandler.GetMovies(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response []models.Movie
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, movie, response)
}

func TestGetMoviesHandler_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()
	movie := &models.PaginationResponse{
		Page:       1,
		PageSize:   1,
		TotalItems: 1,
		TotalPages: 1,
		Data: models.Movie{
			Title:       "Inception",
			Description: "A mind-bending thriller",
			Genre:       "Sci-Fi",
			Duration:    120,
			WatchURL:    "http://example.com/watch",
			Artist:      "Christopher Nolan",
		},
	}

	page := 1
	limit := 5
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().GetMovies(gomock.Any(), gomock.Any()).Return(nil)

	err := movieHandler.GetMovies(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.PaginationResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, movie, response)
}
