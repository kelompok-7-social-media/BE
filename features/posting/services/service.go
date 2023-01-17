package services

import (
	"project/features/posting"

	"github.com/go-playground/validator/v10"
)

type postingSrv struct {
	data     posting.PostingData
	validasi *validator.Validate
}

// Delete implements book.BookService

// Update implements book.BookService

func New(pd posting.PostingData) posting.PostingService {
	return &postingSrv{
		data:     pd,
		validasi: validator.New(),
	}
}

// Add implements posting.PostingService
func (ps *postingSrv) Add(token interface{}, newBook posting.Core) (posting.Core, error) {
	panic("unimplemented")
}
