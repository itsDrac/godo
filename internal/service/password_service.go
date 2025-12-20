package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type passwordService struct {
	cost int
}

func NewPasswordService(cost int) PasswordService {
	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}
	return &passwordService{cost: cost}
}

func (s *passwordService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (s *passwordService) Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
