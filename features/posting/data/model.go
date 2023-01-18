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
	return Posting{
		Model:     gorm.Model{ID: data.ID},
		Postingan: data.Postingan,
		Image_url: data.Image_url,
	}
}
func ToCore(data Posting) posting.Core {
	return posting.Core{
		ID:        data.ID,
		Postingan: data.Postingan,
		Image_url: data.Image_url,
	}
}

func DataToCore(data Posting) posting.Core {
	return posting.Core{
		ID:        data.ID,
		Postingan: data.Postingan,
		Image_url: data.Image_url,
		UserName:  data.Image_url,
		CreatedAt: data.CreatedAt,
	}
}
func DataToCoreArr(data []Posting) []posting.Core {
	var postArr = []posting.Core{}
	for _, listOfPost := range data {
		postArr = append(postArr, DataToCore(listOfPost))
	}

	return postArr
}

// func ToCoreSlice(data []posting.Core) []Image {
// 	images := []Image{}
// 	for _, v := range data {
// 		image := Image{}
// 		image.PostID = v.ID
// 		image.Image_url = v.Image_url

// 		images = append(images, image)
// 	}

// 	return images

// }
