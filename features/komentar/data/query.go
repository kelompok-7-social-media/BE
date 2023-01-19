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
func (kd *komentarData) GetCommentsByPost(userID int, postID int) ([]komentar.Core, error) {
	comments := []KomenUser{}
	if err := kd.db.Where("user_id an posting_id = ?", userID, postID).Find(&comments).Error; err != nil {
		log.Println("Get Mypost query error", err.Error())
		return []komentar.Core{}, err
	}

	var dataCore = ListModelTOCore(comments)

	return dataCore, nil
}

func (kd *komentarData) Delete(userID int, postID int, commentID int) error {
	var record Komentar
	err := kd.db.Where("id = ? AND user_id = ? AND post_id = ? ", commentID, userID, postID).Delete(&record).Error
	if err != nil {
		log.Println("delete comment query error :", err.Error())
		return err
	}
	return nil
}

// Update implements komentar.KomentarData
func (kd *komentarData) Update(userID int, commentID int, updatedData komentar.Core) (komentar.Core, error) {
	cnv := CoreToData(updatedData)
	cnv.ID = uint(commentID)
	cnv.UserID = uint(userID)

	// DB Update(value)
	tx := kd.db.Model(&cnv).Where("user_id = ?", userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update comment query error :", tx.Error)
		return komentar.Core{}, tx.Error

	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update comment query error : data not found")
		return komentar.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}

// GetAllKomen implements komentar.KomentarData
func (kd *komentarData) GetAllKomen() ([]komentar.Core, error) {
	var komentar []KomenUser
	tx := kd.db.Raw("SELECT komentars.id, komentars.comment, komentars.posting_id, users.username FROM komentars JOIN users ON users.id = komentars.user_id JOIN postings ON postings.id = komentars.posting_id WHERE komentars.deleted_at IS NULL").Find(&komentar)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ListModelTOCore(komentar)

	return dataCore, nil
}
