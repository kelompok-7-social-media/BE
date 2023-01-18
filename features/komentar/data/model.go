package data

import (
	"project/features/komentar"

	"gorm.io/gorm"
)

type Komentar struct {
	gorm.Model
	Pesan     string
	PostingID uint
	UserID    uint
}

// type PostUser struct {
// 	ID        uint
// 	Postingan string
// 	UserName  string
// 	Image_url string
// 	CreatedAt time.Time
// }

func CoreToData(data komentar.Core) Komentar {
	return Komentar{
		Model:     gorm.Model{ID: data.ID},
		Pesan:     data.Pesan,
		PostingID: data.PostingID,
	}
}

// func ToCore(data Posting) posting.Core {
// 	return posting.Core{
// 		ID:        data.ID,
// 		Postingan: data.Postingan,
// 		Image_url: data.Image_url,
// 	}
// }

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

// 	return postArr
// }
