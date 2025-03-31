package unit_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

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
	mockProducer := mocks.NewMockProducer(ctrl)

	ratingService := services.NewRatingService(mockRepo, mockProducer)

	// Mock input data
	rating := &models.Rating{
		ID:        utils.GenerateUUID(),
		MovieID:   "movie-123",
		UserID:    "user-456",
		Score:     8.5,
		Comment:   "Great movie!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().RateMovie(gomock.Any(), rating).Return(nil)

	expectedEvent := &models.Event{
		ID:        rating.ID,
		UserID:    rating.UserID,
		MovieID:   rating.MovieID,
		Score:     rating.Score,
		Comment:   rating.Comment,
		CreatedAt: rating.CreatedAt.String(),
		UpdatedAt: rating.UpdatedAt.String(),
	}
	expectedEventData, _ := json.Marshal(expectedEvent)

	mockProducer.EXPECT().ProduceMessage(gomock.Any(), expectedEventData).Return(nil)

	err := ratingService.RateMovie(context.Background(), rating)

	// Assert
	assert.NoError(t, err)
}

func TestRateMovie_Fail_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRatingRepository(ctrl)
	mockProducer := mocks.NewMockProducer(ctrl)

	ratingService := services.NewRatingService(mockRepo, mockProducer)

	rating := &models.Rating{
		ID:      utils.GenerateUUID(),
		MovieID: "movie-123",
		UserID:  "user-456",
		Score:   9.0,
		Comment: "Amazing!",
	}

	mockRepo.EXPECT().RateMovie(gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	err := ratingService.RateMovie(context.Background(), rating)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}
