package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type MovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Duration    int    `json:"duration" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	WatchURL    string `json:"watch_url" validate:"required"`
	Artist      string `json:"artist" validate:"required"`
}

type Movie struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Duration    int                   `json:"duration"`
	Genre       string                `json:"genre"`
	WatchURL    string                `json:"watch_url"`
	Artist      string                `json:"artist"`
	Rating      float64               `json:"rating" gorm:"type:decimal(3,1);default:0.0" `
	IsDeleted   soft_delete.DeletedAt `json:"is_deleted" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	CreatedAt   time.Time             `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
}
