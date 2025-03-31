package models

import "time"

type RatingRequest struct {
	MovieID string  `json:"movie_id" validation:"required"`
	UserID  string  `json:"user_id" validation:"required"`
	Score   float64 `json:"score" validation:"required"`
	Comment string  `json:"comment"`
}

type Rating struct {
	ID        string    `json:"id" gorm:"type:varchar(36);default:uuid_generate_v4()"`
	MovieID   string    `json:"movie_id" gorm:"type:varchar(36);not null;uniqueIndex:idx_movie_user"`
	UserID    string    `json:"user_id" gorm:"type:varchar(36);not null;uniqueIndex:idx_movie_user"`
	Score     float64   `json:"score" gorm:"type:decimal(3,1);not null"`
	Comment   string    `json:"comment" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
