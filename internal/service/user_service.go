package service

import (
	"context"
	"fmt"

	"github.com/itsDrac/godo/internal/db"
)

type userService struct {
	queries         *db.Queries
	passwordService PasswordService
}

func NewUserService(queries *db.Queries, passwordService PasswordService) UserService {
	return &userService{
		queries:         queries,
		passwordService: passwordService,
	}
}

func (s *userService) CreateUser(ctx context.Context, params CreateUserParams) error {
	hashedPassword, err := s.passwordService.Hash(params.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	createdUser, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: hashedPassword,
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("User created successfully! ID: %d, Username: %s, Email: %s\n",
		createdUser.ID, createdUser.Username, createdUser.Email)

	return nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*UserWithPassword, error) {
	dbUser, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &UserWithPassword{
		User: User{
			ID:       dbUser.ID,
			Username: dbUser.Username,
			Email:    dbUser.Email,
		},
		PasswordHash: dbUser.PasswordHash,
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (*User, error) {
	dbUser, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
	}, nil
}
