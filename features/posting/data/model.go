package data

import (
	"project/features/komentar/data"
	"project/features/posting"
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Image_url string
	Postingan string
	UserID    uint
	Komentars []data.Komentar
	// Image     []data.Images

}

type PostUser struct {
	ID        uint
	Postingan string
	Username  string
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

func (dataModel *PostUser) ModelsToCore() posting.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return posting.Core{
		ID:        dataModel.ID,
		Image_url: dataModel.Image_url,
		Postingan: dataModel.Postingan,
		Username:  dataModel.Username,
	}
}

func ListModelTOCore(dataModel []PostUser) []posting.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []posting.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
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
