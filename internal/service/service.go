package service

import "context"

type UserService interface {
	CreateUser(context.Context, CreateUserParams) error
	GetUserByEmail(context.Context, string) (*UserWithPassword, error)
	GetUserByID(context.Context, int32) (*User, error)
}

type AuthService interface {
	Login(context.Context, LoginParams) (LoginResult, error)
	ValidateToken(context.Context, string) (*UserClaims, error)
}

type PasswordService interface {
	Hash(password string) (string, error)
	Verify(password, hash string) bool
}

type UserClaims struct {
	UserID int32
	Expiry int64
}

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserWithPassword struct {
	User
	PasswordHash string
}
