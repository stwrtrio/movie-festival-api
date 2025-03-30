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
