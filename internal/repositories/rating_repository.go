package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
)

type RatingRepository interface {
	RateMovie(ctx context.Context, rating *models.Rating) error
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{db: db}
}

func (r *ratingRepository) RateMovie(ctx context.Context, rating *models.Rating) error {
	return r.db.WithContext(ctx).Create(rating).Error
}
