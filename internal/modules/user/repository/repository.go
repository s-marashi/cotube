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

func (r *inMemUserRepository) CreateUser(id string, name string, email string) (*domain.User, error) {
	_, ok := r.idProjection[id]
	if ok {
		return nil, UserIdExist
	}

	_, ok = r.emailProjection[email]
	if ok {
		return nil, UserEmailExist
	}

	u := &domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	r.idProjection[id] = u
	r.emailProjection[email] = u

	return u, nil
}

var (
	UserIdExist    = errors.New("user id already exists")
	UserEmailExist = errors.New("email already exists")
)
