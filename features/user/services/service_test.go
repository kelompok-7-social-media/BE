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
	repo := mocks.NewUserData(t)
	t.Run("Berhasil Login", func(t *testing.T) {
		inputEmail := "rischi@gmail.com"
		hashed, _ := helper.GeneratePassword("rischi12345")
		resData := user.Core{ID: uint(1), Name: "rischi", Email: "rischi@gmail.com", Username: "ryo17", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil)
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "rischi12345")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		inputEmail := "riau@gmail.com"

		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found"))
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "riau12345")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("password tidak sesuai", func(t *testing.T) {
		inputEmail := "rischi@gmail.com"
		hashed, _ := helper.GeneratePassword("rischi12345")
		resData := user.Core{ID: uint(1), Name: "rischi", Email: "rischi@gmail.com", Username: "ryo17", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil)
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "rischi11111")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password tidak sesuai")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}
func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Name: "rischi", Email: "rischi@gmail.com", Username: "ryo17"}
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	// t.Run("data tidak ditemukan", func(t *testing.T) {
	// 	srv := New(repo)
	// 	_, token := helper.GenerateJWT(1)
	// 	res, err := srv.Profile(token)
	// 	assert.Nil(t, err)

	// 	assert.ErrorContains(t, err, "tidak ditemukan")
	// 	assert.Equal(t, uint(0), res.ID)

	// })
}
func TestAllUser(t *testing.T) {
	repo := mocks.NewUserData(t)
	user := []user.Core{{
		ID:       1,
		Name:     "Herdy",
		Email:    "herdy@gmail.com",
		Username: "herdy123",
	}}
	t.Run("Sukses lihat data", func(t *testing.T) {
		repo.On("AllUser", mock.Anything).Return(user, nil).Once()

		srv := New(repo)
		res, err := srv.AllUser()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)

	})
	t.Run("not found", func(t *testing.T) {
		repo.On("AllUser").Return(nil, errors.New("not found")).Once()

		srv := New(repo)

		res, err := srv.AllUser()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})

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
	// repo := mocks.NewUserData(t)

	// })
}
