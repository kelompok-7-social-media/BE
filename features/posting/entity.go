package posting

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Postingan string `validate:"required"`
	Image     string
	UserName  string `validate:"required"`
	Create_at time.Time
}

type PostingHandler interface {
	Add() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// AllBook() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// MyBook() echo.HandlerFunc
}

type PostingService interface {
	Add(token interface{}, newBook Core) (Core, error)
	// Update(token interface{}, bookID int, updatedData Core) (Core, error)
	// AllBook() ([]Core, error)
	// Delete(token interface{}, bookID int) error
	// MyBook(token interface{}) ([]Core, error)
}

type PostingData interface {
	Add(userID int, newPosting Core) (Core, error)
	// Update(userID int, bookID int, updatedData Core) (Core, error)
	// AllBook() ([]Core, error)
	// Delete(userID int, bookID int) error
	// MyBook(userID int) ([]Core, error)
}
