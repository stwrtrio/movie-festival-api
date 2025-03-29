package models

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type User struct {
	ID           string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"password_hash"`
}
