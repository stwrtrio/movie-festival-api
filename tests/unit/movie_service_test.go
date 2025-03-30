package unit

import (
	"context"
	"errors"
	"testing"

	"github.com/stwrtrio/movie-festival-api/internal/mocks"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovie_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMovieRepository(ctrl)
	movieService := services.NewMovieService(mockRepo)

	movie := &models.Movie{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Rating:      8.8,
	}

	mockRepo.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := movieService.CreateMovie(context.Background(), movie)

	assert.NoError(t, err)
}

func TestCreateMovie_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMovieRepository(ctrl)
	movieService := services.NewMovieService(mockRepo)

	movie := &models.Movie{
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Rating:      8.8,
	}

	mockRepo.EXPECT().CreateMovie(gomock.Any(), gomock.Any()).Return(errors.New("failed to create movie"))

	err := movieService.CreateMovie(context.Background(), movie)

	assert.Error(t, err)
	assert.Equal(t, "failed to create movie", err.Error())
}

func TestUpdateMovie_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMovieRepository(ctrl)
	movieService := services.NewMovieService(mockRepo)

	movie := &models.Movie{
		ID:          "123",
		Title:       "Updated Title",
		Description: "Updated Description",
		Genre:       "Action",
		Rating:      8.5,
	}

	mockRepo.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(nil)

	err := movieService.UpdateMovie(context.Background(), movie)
	assert.NoError(t, err)
}

func TestUpdateMovie_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMovieRepository(ctrl)
	movieService := services.NewMovieService(mockRepo)

	movie := &models.Movie{
		ID:          "123",
		Title:       "Updated Title",
		Description: "Updated Description",
		Genre:       "Action",
		Rating:      8.5,
	}

	mockRepo.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(errors.New("failed to update movie"))

	err := movieService.UpdateMovie(context.Background(), movie)

	assert.Error(t, err)
	assert.Equal(t, "failed to update movie", err.Error())
}
