package data

import (
	comment "project/features/komentar"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) Add(userID int, newComment comment.Core) (comment.Core, error) {
	cnv := CoreToData(newComment)

	err := cd.db.Create(&cnv).Error
	if err != nil {
		return comment.Core{}, err
	}

	newComment.ID = cnv.ID

	return newComment, nil
}

// func (md *menteeData) AddComment(data mentee.CommentsCore) (mentee.CommentsCore, error) {
// 	var input Comments
// 	input = ToEntityComent(data)
// 	res := md.db.Create(&input)
// 	if res.Error != nil {
// 		return mentee.CommentsCore{}, res.Error
// 	}
// 	cnv := FromEntityComment(input)
// 	return cnv, nil

// }
