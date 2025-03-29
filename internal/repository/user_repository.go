package repository

import (
	"errors"
	"test-case-ndi/internal/domain"
)

// struct untuk userRepository
type userRepository struct {
	users       map[int]*domain.User
	usernameMap map[string]int
}

// constructor untuk userRepository
func NewUserRepository() domain.UserRepository {
	// mock data
	users := map[int]*domain.User{
		1: {ID: 1, Username: "orangpertama", Password: "password123", Balance: 1000.50},
		2: {ID: 2, Username: "orangkedua", Password: "password123", Balance: 2500.75},
	}

	usernameMap := make(map[string]int)
	for id, user := range users {
		usernameMap[user.Username] = id
	}

	return &userRepository{
		users:       users,
		usernameMap: usernameMap,
	}
}

// implementasi method GetByID
func (r *userRepository) GetByID(id int) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// implementasi method GetByUsername
func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	id, exists := r.usernameMap[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return r.GetByID(id)
}

// implementasi method GetAll
func (r *userRepository) GetAll() ([]*domain.User, error) {
	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
