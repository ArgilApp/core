package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
	PasswordHash string
}
