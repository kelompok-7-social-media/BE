package data

import (
	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Postingan string
	UserID    uint
	ImageID   uint
}

type BookPemilik struct {
	ID        uint
	Postingan string
	UserName  string
	Image     string
}
