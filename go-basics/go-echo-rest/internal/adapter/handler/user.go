package handler

import (
	"github.com/labstack/echo/v5"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Register(c *echo.Context) error {
	return nil
}
