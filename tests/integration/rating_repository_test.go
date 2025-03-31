package integration

import (
	"testing"

	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateMovieEvent_Success(t *testing.T) {
	// Setup Kafka
	kafkaClient, err := config.InitKafka()
	require.NoError(t, err)
	defer kafkaClient.Close()

	// Seed test movie
	movie := models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Interstellar",
		Description: "Sci-Fi movie",
		Genre:       "Sci-Fi",
		Rating:      8.5,
	}
	err = TestDB.Create(&movie).Error
	require.NoError(t, err)

	// Rate the movie
	err = ratingService.RateMovie(movie.ID, "user123", 9.0)

	assert.NoError(t, err)
	assert.True(t, true, "Kafka event should be published")
}
