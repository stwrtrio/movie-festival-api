package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/mocks"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"
	"gorm.io/gorm"

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
	data := &models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Duration:    120,
		WatchURL:    "http://example.com/watch",
		Artist:      "Christopher Nolan",
		IsDeleted:   0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	movie := &models.PaginationResponse{
		Page:       1,
		PageSize:   5,
		TotalItems: 1,
		TotalPages: 1,
		Data:       data,
	}

	page := 1
	pageSize := 5
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page="+strconv.Itoa(page)+"&page_size="+strconv.Itoa(pageSize), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().GetMovies(gomock.Any(), models.PaginationRequest{Page: page, PageSize: pageSize}).Return(movie, nil)

	err := movieHandler.GetMovies(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedJSON, _ := json.Marshal(movie)
	actualJSON := rec.Body.Bytes()

	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}

func TestGetMoviesHandler_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMovieService(ctrl)
	movieHandler := handlers.NewMovieHandler(mockService)

	e := echo.New()

	page := 1
	pageSize := 5
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page="+strconv.Itoa(page)+"&page_size="+strconv.Itoa(pageSize), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().GetMovies(gomock.Any(), models.PaginationRequest{Page: page, PageSize: pageSize}).Return(nil, errors.New("database error"))

	err := movieHandler.GetMovies(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
