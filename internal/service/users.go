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

// // NewUserService creates a new UserServicer
// func NewUserService(q *db.Queries) *UserServicer {
// 	return &UserServicer{
// 		q: q,
// 	}
// }

// CreateUser creates a new user
func (s *UserServicer) CreateUser(ctx context.Context, u CreateUserParams) error {
	// Implement user creation logic here
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	createdUser, err := s.q.CreateUser(ctx, db.CreateUserParams{
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: hashedPassword,
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Printf("User created successfully! ID: %d, Username: %s, Email: %s\n",
		createdUser.ID, createdUser.Username, createdUser.Email)

	return nil
}

// func (s *UserServicer) Login(ctx context.Context, p LoginParams) (LoginResult, error) {
// 	var out LoginResult

// 	user, err := s.q.GetUserByEmail(ctx, p.Email)
// 	if err != nil {
// 		return out, fmt.Errorf("failed to fetch user: %w", err)
// 	}

// 	if !VerifyPassword(p.Password, user.PasswordHash) {
// 		return out, fmt.Errorf("invalid credentials")
// 	}

// 	secret := utils.GetEnv("JWT_SECRET", "dev-secret")
// 	expiry := time.Now().Add(24 * time.Hour).Unix()
// 	payload := fmt.Sprintf("%d:%d", user.ID, expiry)
// 	mac := hmac.New(sha256.New, []byte(secret))
// 	mac.Write([]byte(payload))
// 	sig := hex.EncodeToString(mac.Sum(nil))
// 	tokStr := fmt.Sprintf("%s.%s", payload, sig)

// 	out = LoginResult{
// 		Token:    tokStr,
// 		UserID:   user.ID,
// 		Username: user.Username,
// 		Email:    user.Email,
// 	}

// 	return out, nil
// }
