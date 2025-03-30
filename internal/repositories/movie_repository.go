package repositories

import (
	"context"
	"errors"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return errors.New("not implemented")
}
