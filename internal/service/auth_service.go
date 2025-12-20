package service

import (
	"context"
	"fmt"

	"github.com/itsDrac/godo/internal/tokens"
)

type authService struct {
	userService     UserService
	passwordService PasswordService
	tokenizer       tokens.Tokenizer
}

func NewAuthService(
	userService UserService,
	passwordService PasswordService,
	tokenizer tokens.Tokenizer,
) AuthService {
	return &authService{
		userService:     userService,
		passwordService: passwordService,
		tokenizer:       tokenizer,
	}
}

func (s *authService) Login(ctx context.Context, params LoginParams) (LoginResult, error) {
	var result LoginResult

	user, err := s.userService.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return result, fmt.Errorf("invalid credentials")
	}

	if !s.passwordService.Verify(params.Password, user.PasswordHash) {
		return result, fmt.Errorf("invalid credentials")
	}

	tokenData := tokens.TokenData{
		UserID: user.ID,
	}

	token, err := s.tokenizer.GenerateToken(tokenData)
	if err != nil {
		return result, fmt.Errorf("failed to generate token: %w", err)
	}

	result = LoginResult{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return result, nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*UserClaims, error) {
	data, err := s.tokenizer.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	tokenData, ok := data.(tokens.TokenData)
	if !ok {
		return nil, fmt.Errorf("invalid token data type")
	}

	return &UserClaims{
		UserID: tokenData.UserID,
		Expiry: tokenData.Expiry,
	}, nil
}
