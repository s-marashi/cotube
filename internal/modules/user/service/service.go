package service

import (
	"github.com/google/uuid"
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
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return s.repository.CreateUser(id.String(), name, email)
}
