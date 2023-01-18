package data

import (
	comment "project/features/komentar"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment   string
	PostingID uint
	UserID    uint
}

// type User struct {
//     gorm.Model
//     Name     string
//     UserName string
//     Comment  []Comment
// }

// type Comment struct {
// 	gorm.Model
// 	Comment   string
// 	PostingID uint
// 	UserID    uint
// }

func ToCore(data Comment) comment.Core {
	return comment.Core{

		ID: data.ID,
		// Posting_ID: data.Posting_ID,
		// User_ID:   data.User_ID,
		Comment: data.Comment,
	}
}

func (dataModel *Comment) ModelsToCore() comment.Core {
	return comment.Core{

		ID: dataModel.ID,
		// Posting_ID: dataModel.Posting_ID,
		// User_ID:   dataModel.User_ID,
		Comment: dataModel.Comment,
	}
}
func listModelToCore(dataModel []Comment) []comment.Core {
	var dataCore []comment.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}
func CoreToData(data comment.Core) Comment {
	return Comment{
		Model: gorm.Model{ID: data.ID},
		// Posting_ID: data.Posting_ID,
		// User_ID:   data.User_ID,
		Comment: data.Comment,
	}
}
