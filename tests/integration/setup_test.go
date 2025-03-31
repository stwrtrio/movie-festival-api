package integration

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var TestDB *gorm.DB
var movieRepo repositories.MovieRepository

func InitTestDB() *gorm.DB {
	if TestDB != nil {
		return TestDB
	}

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

	TestDB = db

	movieRepo = repositories.NewMovieRepository(TestDB)
	// ratingRepo = repositories.NewRatingRepository(TestDB)

	log.Println("Connected to test database")

	return TestDB
}

func CloseTestDB() {
	if TestDB != nil {
		sqlDB, _ := TestDB.DB()
		sqlDB.Close()
	}
}

func TestMain(m *testing.M) {
	InitTestDB()

	code := m.Run()

	CloseTestDB()

	os.Exit(code)
}
