package komentar

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Pesan     string `validate:"required"`
	UserID    uint
	PostingID uint
	CreatedAt time.Time
}

type KomentarHandler interface {
	Add() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// GetAllPost() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// MyPost() echo.HandlerFunc
}

type KomentarService interface {
	Add(token interface{}, newKomen Core) (Core, error)
	// Update(token interface{}, posID int, komenID int, updatedData Core) (Core, error)
	// GetAllPost() ([]Core, error)
	// Delete(token interface{}, bookID int) error
	// MyPost(token interface{}) ([]Core, error)
}

type KomentarData interface {
	Add(userID int, newKomen Core) (Core, error)
	// Update(userID int, posID int, komenID int, updatedData Core) (Core, error)
	// GetAllPost() ([]Core, error)
	// Delete(userID int, bookID int) error
	// MyPost(userID int) ([]Core, error)
}
