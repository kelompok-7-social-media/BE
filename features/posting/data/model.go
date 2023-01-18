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
	Image     []Image
}

type Image struct {
	gorm.Model
	PostID    uint
	Image_url string
}

type PostUser struct {
	ID        uint
	Postingan string
	UserName  string
	Image_url string
	CreatedAt time.Time
}

func CoreToData(data posting.Core) Posting {
	// image := []Image{}
	return Posting{
		Model:     gorm.Model{ID: data.ID},
		Postingan: data.Postingan,
	}
}
func coredata(data posting.Core) Image {
	return Image{
		PostID:    data.ID,
		Image_url: data.Image_url,
	}

}
func ToCoreSlice(data []posting.Core) []Image {
	images := []Image{}
	for _, v := range data {
		image := Image{}
		image.PostID = v.ID
		image.Image_url = v.Image_url

		images = append(images, image)
	}

	return images

}
