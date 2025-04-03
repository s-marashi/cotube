package handlers

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/s-marashi/cotube/internal/modules/user"
	"github.com/s-marashi/cotube/internal/modules/user/repository"

	"net/http"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserRestHandler(service user.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

type registerUserRequest struct {
	Name  string
	Email string
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var request registerUserRequest
	if err := c.Bind(&request); err != nil {
		slog.Debug(fmt.Errorf("payload was invalid: %w", err).Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.userService.RegisterUser(request.Name, request.Email)
	if err != nil {
		if errors.Is(err, repository.UserIdExist) || errors.Is(err, repository.UserEmailExist) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "user already exists"})
		}
		slog.Error(fmt.Errorf("[UserHandler]: failed to register user: %w", err).Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusCreated, user)
}
