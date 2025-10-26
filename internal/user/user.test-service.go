package user

import (
	"fmt"
	"sync"
)

type TestService struct {
	mu    sync.RWMutex
	users []User
}

func TestUserService() UserServiceInterface {
	return &TestService{users: make([]User, 0)}
}

func (m *TestService) Create(user *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// append a copy to avoid external mutation
	u := *user
	m.users = append(m.users, u)
	return nil
}

func (m *TestService) GetUserByUsername(username string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for i := range m.users {
		if m.users[i].Username == username {
			u := m.users[i]
			return &u, nil
		}
	}
	return nil, nil
}

func (m *TestService) GetUserByEmail(email string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for i := range m.users {
		if m.users[i].Email == email {
			u := m.users[i]
			return &u, nil
		}
	}
	return nil, nil
}

func (m *TestService) GetAllUsers() ([]User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// populate mock users once if empty
	if len(m.users) == 0 {
		for i := 1; i <= 10000; i++ {
			user := User{
				Username: fmt.Sprintf("user%d", i),
				Email:    fmt.Sprintf("user%d@example.com", i),
				Password: "password123", // mock password
			}
			m.users = append(m.users, user)
		}
	}

	// return a copy to avoid external mutation
	usersCopy := make([]User, len(m.users))
	copy(usersCopy, m.users)
	return usersCopy, nil
}
