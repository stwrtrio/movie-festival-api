package repositories

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RatingRepository interface {
	RateMovie(ctx context.Context, rating *models.Rating) error
	UpdateMovieRating(ctx context.Context, movieID string) error
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

func (r *ratingRepository) UpdateMovieRating(ctx context.Context, movieID string) error {
	var avgRating float64

	err := r.db.Model(&models.Rating{}).
		Select("COALESCE(AVG(score), 0)").
		Where("movie_id = ?", movieID).
		Scan(&avgRating).Error
	if err != nil {
		return err
	}

	return r.db.Model(&models.Movie{}).
		Where("id = ?", movieID).
		Update("rating", avgRating).Error
}
