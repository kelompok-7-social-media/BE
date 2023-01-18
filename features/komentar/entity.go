package komentar

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	Posting_ID uint
	User_ID    uint
	Comment    string `validate:"required"`
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	// Get() echo.HandlerFunc
	// Delete() echo.HandlerFunc

}

type CommentService interface {
	Add(token interface{}, newComment Core) (Core, error)
	// Get() ([]Core, error)
	// Delete(token interface{}, bookID int) error
}

type CommentData interface {
	Add(userID int, newComment Core) (Core, error)
	// Get() ([]Core, error)
	// Delete(userID int, bookID int) error
}
