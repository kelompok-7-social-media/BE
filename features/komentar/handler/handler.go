package handler

import (
	"fmt"
	"log"
	"net/http"
	"project/features/komentar"
	"project/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type komentarHandle struct {
	srv komentar.KomentarService
}

func New(ps komentar.KomentarService) komentar.KomentarHandler {
	return &komentarHandle{
		srv: ps,
	}
}

// Add implements komentar.KomentarHandler
func (kh *komentarHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddKomenRequest{}
		// idParam := c.Param("posting_id")
		// id, _ := strconv.Atoi(idParam)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := kh.srv.Add(c.Get("user"), *cnv)
		fmt.Println(res)
		if err != nil {
			fmt.Println(err)
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		posting := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan Komentar", posting))
	}
}
func (kh *komentarHandle) GetCommentsByPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, _ := strconv.Atoi(c.Param("post_id"))
		result, _ := kh.srv.GetCommentsByPost(postID)

		listRes := ListCommentCoreToCommentsRespon(result)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan  post", listRes))
	}
}
func (kh *komentarHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, _ := strconv.Atoi(c.Param("post_id"))
		commentID, _ := strconv.Atoi(c.Param("comment_id"))

		del := kh.srv.Delete(c.Get("user"), postID, commentID)
		if del != nil {
			return c.JSON(helper.PrintErrorResponse(del.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menghapus postingan"))
	}
}

func (kh *komentarHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddKomenRequest{}
		KomenID, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)
		fmt.Println(cnv)

		res, err := kh.srv.Update(c.Get("user"), KomenID, *cnv)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		posting := ToResponse("update", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses mengubah komentar", posting))
	}

}

// GetAllKomen implements komentar.KomentarHandler
func (kh *komentarHandle) GetAllKomen() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, _ := kh.srv.GetAllKomen()

		listRes := ListCommentCoreToCommentsRespon(result)
		fmt.Println("ini handler", listRes)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan  Komentar", listRes))
	}
}
