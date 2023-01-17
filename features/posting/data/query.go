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

func (*postingData) Add(userID int, newPosting posting.Core) (posting.Core, error) {
	panic("unimplemented")
}
