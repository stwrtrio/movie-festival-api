package unit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stwrtrio/movie-festival-api/internal/mocks"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/services"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRateMovie_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRatingRepository(ctrl)
	ratingService := services.NewRatingService(mockRepo)

	// Mock input data
	rating := &models.Rating{
		ID:      utils.GenerateUUID(),
		MovieID: "movie-123",
		UserID:  "user-456",
		Score:   8.5,
		Comment: "Great movie!",
	}

	mockRepo.EXPECT().CreateRating(gomock.Any(), gomock.Any()).Return(nil)

	err := ratingService.RateMovie(context.Background(), rating)

	// Assert
	assert.NoError(t, err)
}

func TestRateMovie_Fail_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRatingRepository(ctrl)
	ratingService := services.NewRatingService(mockRepo)

	rating := &models.Rating{
		ID:      utils.GenerateUUID(),
		MovieID: "movie-123",
		UserID:  "user-456",
		Score:   9.0,
		Comment: "Amazing!",
	}

	mockRepo.EXPECT().CreateRating(gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	err := ratingService.RateMovie(context.Background(), rating)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}
