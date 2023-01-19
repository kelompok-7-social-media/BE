package services

import (
	"errors"
	"fmt"
	"log"
	"project/features/komentar"
	"project/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type komentarSrv struct {
	data     komentar.KomentarData
	validasi *validator.Validate
}

func New(pd komentar.KomentarData) komentar.KomentarService {
	return &komentarSrv{
		data:     pd,
		validasi: validator.New(),
	}
}

// Add implements komentar.KomentarService
func (ks *komentarSrv) Add(token interface{}, newKomen komentar.Core) (komentar.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return komentar.Core{}, errors.New("user not found")
	}

	err := ks.validasi.Struct(newKomen)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return komentar.Core{}, errors.New("validation error")
	}

	res, err := ks.data.Add(userID, newKomen)
	fmt.Println(res)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Komentar not found"
		} else {
			msg = "internal server error"
		}
		return komentar.Core{}, errors.New(msg)
	}

	return res, nil
}
func (ks *komentarSrv) GetCommentsByPost(postID int) ([]komentar.Core, error) {
	All, err := ks.data.GetCommentsByPost(postID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "coments not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}
	return All, nil
}
func (ks *komentarSrv) Delete(token interface{}, postID int, commentID int) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := ks.data.Delete(userID, postID, commentID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "coments not found"
		} else {
			msg = "internal server error"

		}
		return errors.New(msg)
	}
	return nil
}

// Update implements komentar.KomentarService
func (ks *komentarSrv) Update(token interface{}, commentID int, updatedData komentar.Core) (komentar.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return komentar.Core{}, errors.New("id user not found")
	}
	if validasieror := ks.validasi.Struct(updatedData); validasieror != nil {
		return komentar.Core{}, nil
	}

	res, err := ks.data.Update(userID, commentID, updatedData)
	if err != nil {
		fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "komentar not found"
		} else {
			msg = "internal server error"
		}
		return komentar.Core{}, errors.New(msg)
	}

	return res, nil
}

// GetAllKomen implements komentar.KomentarService
func (ks *komentarSrv) GetAllKomen() ([]komentar.Core, error) {
	All, err := ks.data.GetAllKomen()
	fmt.Println("ini service", All)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "komentar not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	return All, nil
}
