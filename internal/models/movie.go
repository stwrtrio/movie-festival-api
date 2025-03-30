package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Movie struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Duration    int                   `json:"duration"`
	Genre       string                `json:"genre"`
	WatchURL    string                `json:"watch_url"`
	Artist      string                `json:"artist"`
	Rating      float64               `json:"rating"`
	IsDeleted   soft_delete.DeletedAt `json:"is_deleted" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   gorm.DeletedAt        `json:"deleted_at"`
}
