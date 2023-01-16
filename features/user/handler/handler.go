package handler

import (
	"fmt"
	"net/http"
	helper "project/HELPER"
	"project/features/user"

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

func (uc *userControll) AllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := uc.srv.AllUser()
		if err != nil {
			return c.JSON((helper.PrintErrorResponse(err.Error())))
		}

		dataResp := fromCoreList(result)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan data user", dataResp))
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil login", res, token))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", res))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", res))
	}
}

// // Update implements user.UserHandler
func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("=========inputuser========")
		ex := c.Get("user")
		fmt.Println("=========inputuser2========")
		// id, _ := strconv.Atoi(c.Param("id"))
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		dataCore := *ToCore(input)
		fmt.Println("=========input========")
		res, err := uc.srv.Update(ex, dataCore)
		fmt.Println("=========input2========")
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mengubah data", res))
	}
}

func (uc *userControll) Update2() echo.HandlerFunc {
	return func(c echo.Context) error {
		ex := c.Get("user")
		// id, _ := strconv.Atoi(c.Param("id"))
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		dataCore := *ToCore(input)
		res, err := uc.srv.Update2(ex, dataCore)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mengubah data", res))
	}
}

// // Deactive implements user.UserHandler
func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("user")
		res, err := uc.srv.Delete(tx)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := ToResponse(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil hapus", result))
	}
}
