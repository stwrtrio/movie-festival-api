package models

import (
	"time"
)

type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Genre       string    `json:"genre"`
	WatchURL    string    `json:"watch_url"`
	Artist      string    `json:"artist"`
	Rating      float64   `json:"rating"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
