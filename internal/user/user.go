package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-" gorm:"not null"`
}

type UserServiceInterface interface {
	Create(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
}
