package models

type UserRequest struct {
	Email    string `json:"email,required" gorm:"unique"`
	Password string `json:"password,required"`
}
