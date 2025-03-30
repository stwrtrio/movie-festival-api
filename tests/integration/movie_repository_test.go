package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var movieRepo repositories.MovieRepository

func initTestDB() *gorm.DB {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	testDB = db

	// Init repository
	movieRepo = repositories.NewMovieRepository(testDB)

	return testDB
}

func TestMain(m *testing.M) {
	initTestDB()

	code := m.Run()

	os.Exit(code)
}

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
	testDB.Where("id = ?", movie.ID).Delete(&models.Movie{})
	testDB.Unscoped().Delete(&models.Movie{})
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
	testDB.Create(&movie)

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
	testDB.Where("id = ?", movie.ID).First(&result)

	assert.Equal(t, updatedMovie.Title, result.Title)
	assert.Equal(t, updatedMovie.Description, result.Description)
	assert.Equal(t, updatedMovie.Genre, result.Genre)
	assert.Equal(t, updatedMovie.Rating, result.Rating)

	// Clean up record
	testDB.Where("id = ?", result.ID).Delete(&models.Movie{})
	testDB.Unscoped().Delete(&models.Movie{})
}
