package integration

import (
	"context"
	"testing"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateMovie_Success(t *testing.T) {
	movie := &models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Genre:       "Sci-Fi",
		Rating:      8.8,
	}

	err := movieRepo.CreateMovie(context.Background(), movie)

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, movie.ID)

	// Clean up record
	TestDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", movie.ID).Delete(&models.Movie{})
}

func TestUpdateMovie_Success(t *testing.T) {
	// Insert dummy movie
	movie := &models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Old Title",
		Description: "Old Description",
		Genre:       "Action",
		Rating:      7.0,
	}
	TestDB.Create(&movie)

	// Update movie
	updatedMovie := models.Movie{
		ID:          movie.ID,
		Title:       "New Title",
		Description: "New Description",
		Genre:       "Sci-Fi",
		Rating:      8.5,
	}

	err := movieRepo.UpdateMovie(context.Background(), &updatedMovie)
	assert.NoError(t, err)

	// Verify update
	var result models.Movie
	TestDB.Where("id = ?", movie.ID).First(&result)

	assert.Equal(t, updatedMovie.Title, result.Title)
	assert.Equal(t, updatedMovie.Description, result.Description)
	assert.Equal(t, updatedMovie.Genre, result.Genre)
	assert.Equal(t, updatedMovie.Rating, result.Rating)

	// Clean up record
	TestDB.Where("id = ?", result.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", result.ID).Delete(&models.Movie{})
}

func TestGetMovies_Success(t *testing.T) {
	// Insert dummy movie
	movie := &models.Movie{
		ID:          utils.GenerateUUID(),
		Title:       "Old Title",
		Description: "Old Description",
		Genre:       "Action",
		Rating:      7.0,
	}
	TestDB.Create(&movie)

	movies, err := movieRepo.GetMovies(context.Background(), 1, 1)
	assert.NoError(t, err)
	assert.Len(t, movies, 1)
	assert.Equal(t, movies[0].Title, movie.Title)

	// Clean up record
	TestDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	TestDB.Unscoped().Where("id = ?", movie.ID).Delete(&models.Movie{})
}
