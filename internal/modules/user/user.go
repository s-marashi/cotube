package user

import "github.com/s-marashi/cotube/internal/domain"

type UserService interface {
	RegisterUser(name string, email string) (*domain.User, error)
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
