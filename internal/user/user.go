package user

import "time"

type User struct {
	ID        int        `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type UserServiceInterface interface {
	Create(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
}
