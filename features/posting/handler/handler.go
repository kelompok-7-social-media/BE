package handler

import (
	"log"
	"net/http"
	"project/features/posting"
	"project/helper"

	"github.com/labstack/echo/v4"
)

type postingHandle struct {
	srv posting.PostingService
}

// All implements book.BookHandler

func New(ps posting.PostingService) posting.PostingHandler {
	return &postingHandle{
		srv: ps,
	}
}

func (ph *postingHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdatePostingRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := ph.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		posting := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan posting", posting))
	}
}
func (ph *postingHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error { return c.JSON(http.StatusBadRequest, "format inputan salah") }

}
func (ph *postingHandle) GetAllPost() echo.HandlerFunc {
	return func(c echo.Context) error { return c.JSON(http.StatusBadRequest, "format inputan salah") }
}
func (ph *postingHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error { return c.JSON(http.StatusBadRequest, "format inputan salah") }
}
func (ph *postingHandle) MyPost() echo.HandlerFunc {
	return func(c echo.Context) error { return c.JSON(http.StatusBadRequest, "format inputan salah") }
}
