package handlers

import (
	"colab-tube/internal/domain"
	"colab-tube/internal/modules/user"
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
