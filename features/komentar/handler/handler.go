package handler

import (
	"log"
	"net/http"
	comment "project/features/komentar"
	"project/helper"

	"github.com/labstack/echo/v4"
)

type commentHandle struct {
	srv comment.CommentService
}

// All implements book.BookHandler

func New(cs comment.CommentService) comment.CommentHandler {
	return &commentHandle{
		srv: cs,
	}
}
func (ch *commentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCommentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ToCore(input)

		res, err := ch.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		posting := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan comment", posting))
	}
}

// 		var comment CommentFormat
// 		id_status := c.Param("id")

// 		if err := c.Bind(&comment); err != nil {
// 			c.JSON(http.StatusBadRequest, errors.New("Invalid Input From Client"))
// 		}

// 		idUser, _, role := middlewares.ExtractToken(c)
// 		idCnv, _ := strconv.Atoi(id_status)
// 		idStatus := uint(idCnv)
// 		comment.IdStatus = idStatus
// 		comment.ID_User = uint(idUser)
// 		data := ToDomainComments(comment)
// 		if role == "admin" {
// 			return c.JSON(http.StatusBadRequest, FailedResponse("Invalid Input From Client"))
// 		}
// 		res, err1 := md.MenteeUsecase.Insert(data)
// 		if err1 != nil {
// 			log.Print(err1)
// 			if strings.Contains(err1.Error(), "not found") {
// 				return c.JSON(http.StatusBadRequest, FailedResponse("Not Found"))
// 			}
// 			return c.JSON(http.StatusBadRequest, FailedResponse("Not Found"))
// 		}
// 		return c.JSON(http.StatusCreated, SuccessResponse("success insert comment", ToResponseComments(res)))
