package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
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
