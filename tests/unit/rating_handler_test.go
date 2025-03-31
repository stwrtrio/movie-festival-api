package unit_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/internal/mocks"
	"github.com/stwrtrio/movie-festival-api/internal/models"

	"github.com/golang/mock/gomock"
)

func TestRateMovieHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockRatingService(ctrl)
	ratingHandler := handlers.NewRatingHandler(mockService)

	e := echo.New()
	rating := &models.RatingRequest{
		MovieID: "movie-123",
		UserID:  "user-456",
		Score:   8.5,
		Comment: "Great movie!",
	}

	jsonRating, _ := json.Marshal(rating)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies/rate", bytes.NewReader(jsonRating))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.EXPECT().RateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := ratingHandler.RateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRateMovieHandler_Failed_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockRatingService(ctrl)
	ratingHandler := handlers.NewRatingHandler(mockService)

	e := echo.New()
	rating := &models.RatingRequest{
		MovieID: "movie-123",
		UserID:  "user-456",
		Score:   8.5,
		Comment: "Great movie!",
	}

	jsonRating, _ := json.Marshal(rating)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies/rate", bytes.NewReader(jsonRating))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ratingHandler.RateMovie(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
