package user

import (
	"go-fiber-app/pkg/db"
	"time"
)

type Service struct{}

func UserService() UserServiceInterface {
	return &Service{}
}

func (s *Service) Create(user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `INSERT INTO users (username, email, password, created_at, updated_at) 
	         VALUES (?, ?, ?, ?, ?)`

	result, err := db.GetDB().Exec(query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	// Get the inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (s *Service) GetUserByUsername(username string) (*User, error) {
	var u User
	query := `SELECT id, username, email, password, created_at, updated_at, deleted_at 
	         FROM users WHERE username = ? AND deleted_at IS NULL`

	err := db.GetDB().Get(&u, query, username)
	return &u, err
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	var u User
	query := `SELECT id, username, email, password, created_at, updated_at, deleted_at 
	         FROM users WHERE email = ? AND deleted_at IS NULL`

	err := db.GetDB().Get(&u, query, email)
	return &u, err
}

func (s *Service) GetAllUsers() ([]User, error) {
	var users []User
	query := `SELECT id, username, email, password, created_at, updated_at, deleted_at 
	         FROM users WHERE deleted_at IS NULL`

	err := db.GetDB().Select(&users, query)
	return users, err
}
