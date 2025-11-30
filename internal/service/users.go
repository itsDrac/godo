package service

import (
	"context"
	"fmt"

	"github.com/itsDrac/godo/internal/db"
)

type UserServicer struct {
	// Add dependencies like database connections here
	q *db.Queries
}

// NewUserService creates a new UserServicer
func NewUserService(q *db.Queries) *UserServicer {
	return &UserServicer{
		q: q,
	}
}

// CreateUser creates a new user
func (s *UserServicer) CreateUser(username, email, password string) error {
	// Implement user creation logic here
	if username == "" || email == "" || password == "" {
		return fmt.Errorf("username, email and password are required")
	}

	ctx := context.Background()

	createdUser, err := s.q.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		//PasswordHash: '',
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("User created successfully! ID: %d, Username: %s, Email: %s\n",
		createdUser.ID, createdUser.Username, createdUser.Email)

	return nil
}
