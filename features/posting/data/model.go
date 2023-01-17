package data

import (
	"project/features/posting"
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Postingan string
	UserID    uint
	// Image     []data.Images
}

type PostUser struct {
	ID        uint
	Postingan string
	UserName  string
	Image_url string
	CreateAt  time.Time
}

func CoreToData(data posting.Core) Posting {
	return Posting{
		Model:     gorm.Model{ID: data.ID},
		Postingan: data.Postingan,
	}
}
