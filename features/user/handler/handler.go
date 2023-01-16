package handler

import (
	"api/features/user"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login(c echo.Context) error {
	return func(c echo.Context) error {

	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}

// Update implements user.UserHandler
func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}

// Deactive implements user.UserHandler
func (uc *userControll) Deactive() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
