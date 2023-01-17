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

type PostUser struct {
	ID        uint
	Postingan string
	UserName  string
	Image     string
}
