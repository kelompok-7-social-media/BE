package komentar

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Comment   string `validate:"required"`
	PostingID uint
	Username  string
	UserID    uint
	CreatedAt string
}

type KomentarHandler interface {
	Add() echo.HandlerFunc
	GetCommentsByPost() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAllKomen() echo.HandlerFunc
}

type KomentarService interface {
	Add(token interface{}, newComment Core) (Core, error)
	GetCommentsByPost(postID int) ([]Core, error)
	Delete(token interface{}, commentID int) error
	Update(token interface{}, commentID int, updatedComment Core) (Core, error)
	GetAllKomen() ([]Core, error)
}

type KomentarData interface {
	Add(userID int, newComment Core) (Core, error)
	GetCommentsByPost(postID int) ([]Core, error)
	Delete(UserID int, commentID int) error
	Update(UserID, commentID int, updatedComment Core) (Core, error)
	GetAllKomen() ([]Core, error)
}
