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

// Update implements komentar.KomentarService
func (ks *komentarSrv) Update(token interface{}, komenID int, postID int, updatedData komentar.Core) (komentar.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return komentar.Core{}, errors.New("id user not found")
	}
	if validasieror := ks.validasi.Struct(updatedData); validasieror != nil {
		return komentar.Core{}, nil
	}

	res, err := ks.data.Update(userID, komenID, postID, updatedData)
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
