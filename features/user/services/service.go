package services

import (
	"errors"
	"fmt"
	"log"
	helper "project/HELPER"
	user "project/features/user"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) AllUser() ([]user.Core, error) {
	data, err := uuc.qry.AllUser()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "terdapat masalah pada server"

		}
		return nil, errors.New(msg)

	}
	return data, nil
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	token, _ := helper.GenerateJWT(int(res.ID))

	return token, res, nil

}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	hashed, err := helper.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(token interface{}, updateData user.Core) (user.Core, error) {
	hashed, err := helper.GeneratePassword(updateData.Password)
	id := helper.ExtractToken(token)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	updateData.Password = string(hashed)
	fmt.Println("========password========")
	res, err := uuc.qry.Update(uint(id), updateData)
	fmt.Println("========password2========")
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Update2(token interface{}, updateData user.Core) (user.Core, error) {
	hashed, err := helper.GeneratePassword(updateData.Password)
	id := helper.ExtractToken(token)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	updateData.Password = string(hashed)
	fmt.Println("========password========")
	res, err := uuc.qry.Update(uint(id), updateData)
	fmt.Println("========password2========")
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}
func (uuc *userUseCase) Delete(token interface{}) (user.Core, error) {

	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("id user not found")
	}
	data, err := uuc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return data, nil

}
