package service

import "context"

type Servicer interface {
	CreateUser(context.Context, CreateUserParams) error
}