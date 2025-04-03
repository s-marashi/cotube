package repository

import (
	"errors"

	"github.com/s-marashi/cotube/internal/domain"
	"github.com/s-marashi/cotube/internal/modules/user"
)

type inMemUserRepository struct {
	idProjection    map[string]*domain.User
	emailProjection map[string]*domain.User
}

func NewInMemUserRepository() user.UserRepository {
	return &inMemUserRepository{
		idProjection:    make(map[string]*domain.User),
		emailProjection: make(map[string]*domain.User),
	}
}

func (r *inMemUserRepository) CreateUser(user *domain.User) error {
	_, ok := r.idProjection[user.Id()]
	if ok {
		return UserIdExist
	}

	_, ok = r.emailProjection[user.Email()]
	if ok {
		return UserEmailExist
	}

	r.idProjection[user.Id()] = user
	r.emailProjection[user.Email()] = user

	return nil
}

func (r *inMemUserRepository) FindByEmail(email string) (*domain.User, error) {
	user, ok := r.emailProjection[email]
	if ! ok {
		return nil, UserNotFound
	}

	return user, nil
}

var (
	UserIdExist    = errors.New("user id already exists")
	UserEmailExist = errors.New("email already exists")
	UserNotFound = errors.New("user not found")
)
