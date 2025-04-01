package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	GetMovies(ctx context.Context, page, limit int) ([]models.Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *movieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	return r.db.WithContext(ctx).Updates(movie).Error
}

func (r *movieRepository) GetMovies(ctx context.Context, page, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (page - 1) * limit

	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&movies).Error

	return movies, err
}
