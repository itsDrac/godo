package service

type Servicer interface {
	CreateUser(username, email, password string) error
}