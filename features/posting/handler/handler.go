package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project/features/posting"
	"project/helper"
	"strconv"

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
		file, errPath := c.FormFile("file")

		fmt.Print("error get path handler, err = ", errPath)

		if file != nil {
			res, err := helper.UploadImage(c)
			// fmt.Println(res)
			if err != nil {
				fmt.Println(err)
				return errors.New("Create gambar Failed. Cannot Upload Data.")
			}
			input.Image_url = res
			// fmt.Println(input.Image_url)
		} else {
			input.Image_url = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := input.reqToCore()

		res, err := ph.srv.Add(c.Get("user"), cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		posting := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan posting", posting))
	}
}
func (ph *postingHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdatePostingRequest{}
		file, errPath := c.FormFile("file")

		fmt.Print("error get path handler, err = ", errPath)

		if file != nil {
			res, err := helper.UploadImage(c)
			// fmt.Println(res)
			if err != nil {
				fmt.Println(err)
				return errors.New("Create gambar Failed. Cannot Upload Data.")
			}
			input.Image_url = res
			// fmt.Println(input.Image_url)
		} else {
			input.Image_url = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)

		PostID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		res, err := ph.srv.Update(c.Get("user"), PostID, *cnv)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		posting := ToResponse("update", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses mengubah buku", posting))
	}

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
