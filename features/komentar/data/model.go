package data

import (
	"project/features/komentar"

	"gorm.io/gorm"
)

type Komentar struct {
	gorm.Model
	Comment   string
	PostingID uint
	UserID    uint
}

type KomenUser struct {
	ID        uint
	Comment   string
	PostingID uint
	Username  string
}

func CoreToData(data komentar.Core) Komentar {
	return Komentar{
		Model:     gorm.Model{ID: data.ID},
		Comment:   data.Comment,
		PostingID: data.PostingID,
		UserID:    data.UserID,
	}
}

func ToCore(data Komentar) komentar.Core {
	return komentar.Core{
		ID:      data.ID,
		Comment: data.Comment,
	}
}
func DataToCoreArr(data []Komentar) []komentar.Core {
	var commentArr = []komentar.Core{}
	for _, listOfComment := range data {
		commentArr = append(commentArr, ToCore(listOfComment))
	}

	return commentArr
}

// func DataToCore(data Posting) posting.Core {
// 	return posting.Core{
// 		ID:        data.ID,
// 		Postingan: data.Postingan,
// 		Image_url: data.Image_url,
// 		UserName:  data.Image_url,
// 		CreatedAt: data.CreatedAt,
// 	}
// }
// func DataToCoreArr(data []Posting) []posting.Core {
// 	var postArr = []posting.Core{}
// 	for _, listOfPost := range data {
// 		postArr = append(postArr, DataToCore(listOfPost))
// 	}

//		return postArr
//	}
func (dataModel *KomenUser) ModelsToCore() komentar.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return komentar.Core{
		ID:        dataModel.ID,
		Comment:   dataModel.Comment,
		PostingID: dataModel.PostingID,
		Username:  dataModel.Username,
	}
}

func ListModelTOCore(dataModel []KomenUser) []komentar.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []komentar.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
