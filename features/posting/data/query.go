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
	res := []Posting{}
	err := pd.db.Find(&res).Error
	if err != nil {
		log.Println("getall post query error : ", err.Error())
		return DataToCoreArr(res), err

	}
	return []posting.Core{}, nil
}
func (pd *postingData) Delete(userID int, postID int) error {
	var record Posting
	err := pd.db.Where("id = ? AND user_id = ?", postID, userID).Delete(&record).Error
	if err != nil {
		log.Println("delete post query error :", err.Error())
		return err
	}
	return nil
}
func (pd *postingData) MyPost(userID int) ([]posting.Core, error) {
	myPost := []Posting{}
	if err := pd.db.Where("user_id = ?", userID).Find(&myPost).Error; err != nil {
		log.Println("Get Mypost query error", err.Error())
		return []posting.Core{}, err
	}
	return DataToCoreArr(myPost), nil
}
