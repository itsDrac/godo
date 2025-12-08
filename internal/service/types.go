package service

import "time"

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty,omitzero"`
}

type CreatedUserOutput struct {
	CreateUserParams
	UserID    string    `json:"user_id,omitempty,omitzero"`
	CreatedAt time.Time `json:"created_at,omitempty,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitempty,omitzero"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token    string `json:"token"`
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
