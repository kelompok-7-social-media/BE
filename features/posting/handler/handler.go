package handler

import (
	"errors"
	"fmt"
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
		file, _ := c.FormFile("file")
		if file != nil {
			res, err := helper.UploadImage(c)
			fmt.Println(res)
			if err != nil {
				return errors.New("Create gambar Failed. Cannot Upload Data.")
			}
			input.Image_url = res
			fmt.Println(input.Image_url)
		} else {
			input.Image_url = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
		}

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
