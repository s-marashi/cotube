package domain

import (
	"github.com/google/uuid"
	"github.com/s-marashi/cotube/pkg/err"
)

// User represents a user in the system.
type User struct {
	id    string
	name  string
	email string
}

func NewUser(name string, email string) (*User, error) {
	if name == "" {
		return nil, err.NewDomainError("user", "User", "name", "cannot be empty")
	}

	if email == "" {
		return nil, err.NewDomainError("user", "User", "email", "cannot be empty")
	}

	id, e := uuid.NewV7()
	if e != nil {
		return nil, err.NewDomainError("user", "User", "id", "failed to generate")
	}

	return &User{
		id:    id.String(),
		name:  name,
		email: email,
	}, nil
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}
