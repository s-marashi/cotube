package user

import "colab-tube/internal/domain"

type UserService interface {
	RegisterUser(name string, email string) (*domain.User, error)
}

type UserRepository interface {

}