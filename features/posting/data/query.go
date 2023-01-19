package data

import (
	"errors"
	"log"
	"project/features/posting"
	"strings"

	"gorm.io/gorm"
)

type postingData struct {
	db *gorm.DB
}

// Add implements posting.PostingData

// Delete implements book.BookData

func New(db *gorm.DB) posting.PostingData {
	return &postingData{
		db: db,
	}
}

func (pd *postingData) Add(userID int, newPosting posting.Core) (posting.Core, error) {
	cnv := CoreToData(newPosting)
	cnv.UserID = uint(userID)

	err := pd.db.Create(&cnv).Error
	if err != nil {
		log.Println("add post query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return posting.Core{}, errors.New(msg)
	}

	newPosting.ID = cnv.ID

	return newPosting, nil
}

func (pd *postingData) Update(userID int, postID int, updatedData posting.Core) (posting.Core, error) {

	cnv := CoreToData(updatedData)

	// DB Update(value)
	tx := pd.db.Where("id = ? AND user_id = ?", postID, userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update book query error :", tx.Error)
		return posting.Core{}, tx.Error

	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return posting.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}

func (pd *postingData) GetAllPost() ([]posting.Core, error) {
	var komentar []PostUser
	tx := pd.db.Raw("SELECT postings.id, postings.postingan, postings.image_url, users.username FROM postings JOIN users ON users.id = postings.user_id  WHERE postings.deleted_at IS NULL").Find(&komentar)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = ListModelTOCore(komentar)

	return dataCore, nil
}
func (pd *postingData) Delete(userID int, postID int) error {
	post := Posting{}
	err := pd.db.Where("id = ? AND user_id = ?", postID, userID).Delete(&post, postID)
	if err.Error != nil {
		log.Println("delete book query error :", err.Error)
		return err.Error
	}
	if err.RowsAffected <= 0 {
		log.Println("delete book query error : data not found")
		return errors.New("not found")
	}

	return nil
}
func (pd *postingData) MyPost(userID int) ([]posting.Core, error) {
	var myBooks []PostUser
	err := pd.db.Raw("SELECT postings.id, postings.postingan, postings.image_url, users.username FROM postings JOIN users ON users.id = postings.user_id WHERE postings.user_id = ?", userID).Find(&myBooks).Error
	if err != nil {
		return nil, err
	}

	var dataCore = ListModelTOCore(myBooks)

	return dataCore, nil
}
