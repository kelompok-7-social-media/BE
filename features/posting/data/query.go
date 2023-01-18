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
	return posting.Core{}, nil
}
func (pd *postingData) GetAllPost() ([]posting.Core, error) {
	return []posting.Core{}, nil
}
func (pd *postingData) Delete(userID int, postID int) error {
	return nil
}
func (pd *postingData) MyPost(userID int) ([]posting.Core, error) {
	return []posting.Core{}, nil
}
