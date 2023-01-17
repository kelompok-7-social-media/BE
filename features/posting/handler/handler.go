package handler

import (
	"project/features/posting"

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
		// input := AddUpdateBookRequest{}
		// if err := c.Bind(&input); err != nil {
		// 	return c.JSON(http.StatusBadRequest, "format inputan salah")
		// }

		// cnv := ToCore(input)

		// res, err := bh.srv.Add(c.Get("user"), *cnv)
		// if err != nil {
		// 	log.Println("trouble :  ", err.Error())
		// 	return c.JSON(helper.PrintErrorResponse(err.Error()))
		// }

		// book := ToResponse("add", res)

		// return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", book))
	}
}
