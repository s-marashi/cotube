package user

import "github.com/s-marashi/cotube/internal/domain"

type UserService interface {
	RegisterUser(name string, email string) (*domain.User, error)
}

type UserRepository interface {
	CreateUser(id string, name string, email string) (*domain.User, error)
}
