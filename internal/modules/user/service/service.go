package service

import (
	"fmt"

	"github.com/s-marashi/cotube/internal/domain"
	"github.com/s-marashi/cotube/internal/modules/user"
)

type userService struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) user.UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) RegisterUser(name string, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, fmt.Errorf("[UserService]: failed to create user: %w", err)
	}

	err = s.repository.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
