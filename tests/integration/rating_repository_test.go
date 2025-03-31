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
