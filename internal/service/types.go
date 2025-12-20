package service

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
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
