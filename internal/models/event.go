package models

type Event struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	MovieID   string  `json:"movie_id"`
	Score     float64 `json:"score"`
	Comment   string  `json:"comment"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
