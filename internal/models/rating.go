package models

import "time"

type Rating struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	MovieID   string    `json:"movie_id" gorm:"type:varchar(36);index;not null"`
	UserID    string    `json:"user_id" gorm:"type:varchar(36);index;not null"`
	Score     float64   `json:"score" gorm:"type:decimal(3,1);not null"`
	Comment   string    `json:"comment" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
