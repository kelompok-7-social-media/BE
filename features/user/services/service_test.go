package services

import (
	"errors"
	"project/features/user"
	helper "project/helper"
	"project/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {

}

func TestProfile(t *testing.T) {

}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("sukses update data", func(t *testing.T) {
		input := user.Core{Name: "herdi", Email: "herdiladania11@gmail.com", Username: "herdiladania11"}
		passwordHashed, _ := helper.GeneratePassword("herdi123")
		updatedData := user.Core{ID: uint(1), Name: "herdi", Email: "herdiladania11@gmail.com", Username: "herdiladania11", Password: passwordHashed}
		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.Nil(t, err)
		assert.Equal(t, updatedData.ID, res.ID)
		assert.NotEqual(t, input.Name, res.Name)
		assert.NotEqual(t, input.Email, res.Email)
		assert.NotEqual(t, input.Username, res.Username)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		input := user.Core{Name: "herdi", Email: "herdiladania11@gmail.com", Username: "herdiladania11"}
		service := New(repo)

		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		input := user.Core{Name: "herdi", Email: "herdiladania11@gmail.com", Username: "herdiladania11"}
		repo.On("Update", uint(5), input).Return(user.Core{}, errors.New("data not found")).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		input := user.Core{Name: "herdi", Email: "herdiladania11@gmail.com", Username: "herdiladania11"}
		repo.On("Update", uint(1), input).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("sukses menghapus profile", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(user.Core{}, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(token)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		_, err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(5)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}
func TestRegister(t *testing.T) {

}
