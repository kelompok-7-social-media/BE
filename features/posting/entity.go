package posting

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Postingan string `validate:"required"`
	Image_url string
	UserName  string
	CreatedAt time.Time
}

type PostingHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAllPost() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyPost() echo.HandlerFunc
}

type PostingService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Update(token interface{}, bookID int, updatedData Core) (Core, error)
	GetAllPost() ([]Core, error)
	Delete(token interface{}, bookID int) error
	MyPost(token interface{}) ([]Core, error)
}

type PostingData interface {
	Add(userID int, newPosting Core) (Core, error)
	Update(userID int, bookID int, updatedData Core) (Core, error)
	GetAllPost() ([]Core, error)
	Delete(userID int, bookID int) error
	MyPost(userID int) ([]Core, error)
}
