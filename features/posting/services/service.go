package services

import (
	"errors"
	"fmt"
	"project/features/posting"
	"project/helper"
	"strings"

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
func (ps *postingSrv) Add(token interface{}, newPosting posting.Core) (posting.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return posting.Core{}, errors.New("user not found")
	}

	// err := ps.validasi.Struct(newPosting)
	// if err != nil {
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		log.Println(err)
	// 	}
	// 	return posting.Core{}, errors.New("validation error")
	// }

	res, err := ps.data.Add(userID, newPosting)
	fmt.Println(res)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Posting not found"
		} else {
			msg = "internal server error"
		}
		return posting.Core{}, errors.New(msg)
	}

	return res, nil
}
func (ps *postingSrv) GetAllPost() ([]posting.Core, error) {
	All, err := ps.data.GetAllPost()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	return All, nil
}

func (ps *postingSrv) Update(token interface{}, postID int, updatedData posting.Core) (posting.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return posting.Core{}, errors.New("id user not found")
	}
	if validasieror := ps.validasi.Struct(updatedData); validasieror != nil {
		return posting.Core{}, nil
	}

	res, err := ps.data.Update(userID, postID, updatedData)
	if err != nil {
		fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Posting not found"
		} else {
			msg = "internal server error"
		}
		return posting.Core{}, errors.New(msg)
	}

	return res, nil

}
func (ps *postingSrv) Delete(token interface{}, postID int) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := ps.data.Delete(userID, postID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"

		}
		return errors.New(msg)
	}
	return nil
}

func (ps *postingSrv) MyPost(token interface{}) ([]posting.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return nil, errors.New("user not found")
	}

	res, _ := ps.data.MyPost(userID)

	return res, nil
}
