package data

import (
	"project/features/komentar/data"
	"project/features/posting"
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Postingan string
	UserID    uint
	Comments  []data.Comment `gorm:"foreignKey:PostingID;references:ID;constraint:OnDelete:CASCADE"`
	// Image     []data.Images
}

type Comment struct {
	PostingID uint
	UserID    uint
	Comment   string
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
