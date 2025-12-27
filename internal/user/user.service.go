package user

import "go-fiber-app/pkg/db"

type Service struct{}

func UserService() UserServiceInterface {
	return &Service{}
}

func (s *Service) Create(user *User) error {
	return db.GetDB().Create(user).Error
}

func (s *Service) GetUserByUsername(username string) (*User, error) {
	var u User
	err := db.GetDB().Where("username = ?", username).First(&u).Error
	return &u, err
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	var u User
	err := db.GetDB().Where("email = ?", email).First(&u).Error
	return &u, err
}

func (s *Service) GetAllUsers() ([]User, error) {
	var users []User
	err := db.GetDB().Find(&users).Error
	return users, err
}
