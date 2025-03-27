package handlers

import (
	"github.com/s-marashi/cotube/internal/domain"
	"github.com/s-marashi/cotube/internal/modules/user"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) RegisterUser(name string, email string) (*domain.User, error) {
	return h.userService.RegisterUser(name, email)
}
