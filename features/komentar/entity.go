package komentar

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Comment   string `validate:"required"`
	UserID    uint
	PostingID uint
	CreatedAt string
}

type KomentarHandler interface {
	Add() echo.HandlerFunc
	GetCommentsByPost() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type KomentarService interface {
	Add(token interface{}, newComment Core) (Core, error)
	GetCommentsByPost(postID int) ([]Core, error)
	Delete(token interface{}, postID int, commentID int) error
}

type KomentarData interface {
	Add(userID int, newComment Core) (Core, error)
	GetCommentsByPost(postID int) ([]Core, error)
	Delete(UserID int, postID int, commentID int) error
}
