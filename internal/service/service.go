package service

import "context"

type Servicer interface {
	CreateUser(context.Context, CreateUserParams) error
	Login(context.Context, LoginParams) (LoginResult, error)
}
