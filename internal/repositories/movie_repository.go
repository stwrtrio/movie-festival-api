package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	GetMovies(ctx context.Context, pagination models.PaginationRequest) ([]models.Movie, int64, error)
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

func (r *movieRepository) GetMovies(ctx context.Context, pagination models.PaginationRequest) ([]models.Movie, int64, error) {
	var movies []models.Movie
	var total int64

	offset := (pagination.Page - 1) * pagination.PageSize

	if err := r.db.WithContext(ctx).Model(&models.Movie{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).
		Limit(pagination.PageSize).
		Offset(offset).
		Find(&movies).Error
	if err != nil {
		return nil, 0, err
	}

	return movies, total, err
}
