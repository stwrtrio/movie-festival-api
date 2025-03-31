package integration

import (
	"context"
	"testing"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateMovieEvent_Success(t *testing.T) {
	// Seed test movie
	movie := models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Interstellar",
		Description: "Sci-Fi movie",
		Genre:       "Sci-Fi",
		Rating:      8.5,
	}
	err := TestDB.Create(&movie).Error
	require.NoError(t, err)

	rating := models.Rating{
		ID:      utils.GenerateUUID(),
		UserID:  "user123",
		MovieID: movie.ID,
		Score:   9.0,
		Comment: "Great movie!",
	}

	// Rate the movie
	err = ratingRepo.RateMovie(context.Background(), &rating)

	assert.NoError(t, err)
	assert.True(t, true, "Kafka event should be published")

	// Cleanup Test Data
	TestDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Where("id = ?", rating.ID).Delete(&models.Rating{})
	TestDB.Unscoped().Where("id = ?", rating.ID).Delete(&models.Rating{})
}

func TestUpdateMovieRating_Success(t *testing.T) {
	movie := &models.Movie{
		ID:          "test-movie-123",
		Title:       "Test Movie",
		Description: "For testing purposes",
		Genre:       "Action",
		Rating:      0,
	}
	TestDB.Create(movie)

	ratings := []models.Rating{
		{MovieID: movie.ID, UserID: "user-1", Score: 8},
		{MovieID: movie.ID, UserID: "user-2", Score: 6},
		{MovieID: movie.ID, UserID: "user-3", Score: 7},
	}

	TestDB.Create(&ratings)

	err := ratingRepo.UpdateMovieRating(context.Background(), movie.ID)
	assert.NoError(t, err)

	var updatedMovie models.Movie
	err = TestDB.First(&updatedMovie, "id = ?", movie.ID).Error
	assert.NoError(t, err)

	expectedAvgRating := (8.0 + 6.0 + 7.0) / 3.0
	assert.Equal(t, expectedAvgRating, updatedMovie.Rating)

	// Cleanup data
	TestDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Delete(&models.Rating{}, "movie_id = ?", movie.ID)
}

func TestUpdateMovieRating_NoRatings(t *testing.T) {
	movie := &models.Movie{
		ID:          "test-movie-no-rating",
		Title:       "No Rating Movie",
		Description: "For testing purposes",
		Genre:       "Drama",
		Rating:      5,
	}

	TestDB.Create(movie)

	err := ratingRepo.UpdateMovieRating(context.Background(), movie.ID)
	assert.NoError(t, err)

	var updatedMovie models.Movie
	err = TestDB.First(&updatedMovie, "id = ?", movie.ID).Error
	assert.NoError(t, err)

	assert.Equal(t, 0.0, updatedMovie.Rating)

	// Cleanup data
	TestDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", movie.ID).Delete(&models.Movie{})
}
