package data

import (
	"errors"
	"log"
	"project/features/komentar"
	"strings"

	"gorm.io/gorm"
)

type komentarData struct {
	db *gorm.DB
}

func New(db *gorm.DB) komentar.KomentarData {
	return &komentarData{
		db: db,
	}
}

// Add implements komentar.KomentarData
func (kd *komentarData) Add(userID int, newKomen komentar.Core) (komentar.Core, error) {
	cnv := CoreToData(newKomen)
	cnv.UserID = uint(userID)

	err := kd.db.Create(&cnv).Error
	if err != nil {
		log.Println("add post query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return komentar.Core{}, errors.New(msg)
	}

	newKomen.ID = cnv.ID

	return newKomen, nil
}
func (kd *komentarData) GetCommentsByPost(postID int) ([]komentar.Core, error) {
	comments := []Komentar{}
	if err := kd.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		log.Println("Get Mypost query error", err.Error())
		return []komentar.Core{}, err
	}
	return DataToCoreArr(comments), nil
}

func (kd *komentarData) Delete(userID, postID int, commentID int) error {
	var record Komentar
	err := kd.db.Where("id = ? AND user_id = ? AND post_id = ? ", commentID, userID, postID).Delete(&record).Error
	if err != nil {
		log.Println("delete comment query error :", err.Error())
		return err
	}
	return nil
}

// Update implements komentar.KomentarData
// func (kd *komentarData) Update(userID int, komenID int, postID int, updatedData komentar.Core) (komentar.Core, error) {
// 	cnv := CoreToData(updatedData)

// 	// DB Update(value)
// 	tx := kd.db.Where("id = ? && user_id = ? && posting_id = ?", komenID, userID, postID).Updates(&cnv)
// 	if tx.Error != nil {
// 		log.Println("update book query error :", tx.Error)
// 		return komentar.Core{}, tx.Error

// 	}

// 	// Rows affected checking
// 	if tx.RowsAffected <= 0 {
// 		log.Println("update book query error : data not found")
// 		return komentar.Core{}, errors.New("not found")
// 	}

// 	// return result converting cnv to book.Core
// 	return ToCore(cnv), nil
// }
