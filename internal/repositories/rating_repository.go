package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	result := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "movie_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"score":      rating.Score,
			"comment":    rating.Comment,
			"updated_at": gorm.Expr("NOW()"),
		}),
	}).Create(&rating)

	return result.Error
}
