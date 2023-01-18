package data

import (
	"project/features/posting"

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
		return posting.Core{}, err
	}

	newPosting.ID = cnv.ID

	return newPosting, nil
}
