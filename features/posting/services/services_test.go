package services

import (
	"errors"
	"project/features/posting"
	"project/helper"
	"project/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewPostingData(t)
	t.Run("Berhasil Menambahkan Postingan", func(t *testing.T) {
		type SamplePost struct {
			ID        uint
			Postingan string
			Image_url string
			Username  string
			CreatedAt time.Time
		}
		inputData := SamplePost{
			ID:        uint(1),
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}
		resData := posting.Core{
			ID:        uint(1),
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}

		Respon := posting.Core{
			ID:        uint(1),
			Postingan: inputData.Postingan,
			Image_url: inputData.Image_url,
			Username:  inputData.Username,
			CreatedAt: inputData.CreatedAt,
		}

		repo.On("Add", mock.Anything).Return(Respon, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, resData)
		assert.Nil(t, err)
		assert.Equal(t, Respon.ID, res.ID)
		assert.Equal(t, resData.Postingan, res.Postingan)
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		inputData := posting.Core{
			ID:        0,
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}
		resData := posting.Core{
			ID:        1,
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}

		repo.On("Add", mock.Anything).Return(inputData, nil)
		srv := New(repo)
		_, token := helper.GenerateJWT(0)
		res, err := srv.Add(token, resData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user not found")
		assert.Empty(t, res)
		assert.Equal(t, uint(0), resData)
		repo.AssertExpectations(t)
	})
	t.Run("masalah diserver", func(t *testing.T) {
		type SamplePost struct {
			ID        uint
			Postingan string
			Image_url string
			Username  string
			CreatedAt time.Time
		}
		inputData := SamplePost{
			ID:        uint(1),
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}
		resData := posting.Core{
			ID:        uint(1),
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
			CreatedAt: time.Time{},
		}
		repo.On("Add", mock.Anything).Return(posting.Core{}, errors.New("internal server error")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(int(inputData.ID))
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, resData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, res)
		assert.Equal(t, uint(0), resData)
		repo.AssertExpectations(t) //new service
	})
}

func TestDeletePost(t *testing.T) {
	repo := mocks.NewPostingData(t)

	srv := New(repo)
	t.Run("Delete Success", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Delete Error", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("user id not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)

		repo.AssertExpectations(t)
	})
	t.Run("Delete Error", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})
	t.Run("Delete server error", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("internal server error")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestMyPost(t *testing.T) {
	repo := mocks.NewPostingData(t)

	srv := New(repo)

	// Case: user ingin melihat list buku yang dimilikinya
	t.Run("Post succesfully", func(t *testing.T) {
		resData := []posting.Core{
			{
				ID:        1,
				Postingan: "Hari yang menyenangkan",
				Image_url: "www.google.com",
				Username:  "ryo17",
			},
			{
				ID:        2,
				Postingan: "Hari yang menyenangkan",
				Image_url: "www.google.com",
				Username:  "ryo17",
			},
		}

		// Programming input and return repo
		repo.On("MyPost", 1).Return(resData, nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.MyPost(token)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, actual[0].ID)
		assert.Equal(t, resData[0].Postingan, actual[0].Postingan)
		assert.Equal(t, resData[1].ID, actual[1].ID)
		assert.Equal(t, resData[1].Postingan, actual[1].Postingan)
	})
}

func TestGetAllPost(t *testing.T) {
	repo := mocks.NewPostingData(t)

	srv := New(repo)
	resData := []posting.Core{
		{
			ID:        1,
			Postingan: "Hari yang menyenangkan",
			Image_url: "www.google.com",
			Username:  "ryo17",
		}}

	// Case: user ingin melihat post
	t.Run("Post succesfully", func(t *testing.T) {

		// Programming input and return repo
		repo.On("AllPost", mock.Anything).Return(resData, nil).Once()

		res, err := srv.GetAllPost()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)

	})
	t.Run("not found", func(t *testing.T) {
		repo.On("AllPost").Return(nil, errors.New("not found")).Once()

		res, err := srv.GetAllPost()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdatePost(t *testing.T) {
	input := posting.Core{Postingan: "One Piece"}
	resData := posting.Core{
		ID:        1,
		Postingan: "Hari yang menyenangkan",
		Image_url: "www.google.com",
		Username:  "ryo17",
	}

	repo := mocks.NewPostingData(t)
	srv := New(repo)

	t.Run("Update successfully", func(t *testing.T) {
		repo.On("Update", 1, 1, input).Return(resData, nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.Update(token, 1, input)

		assert.Nil(t, err)
		assert.Equal(t, resData.Postingan, actual.Postingan)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Username, actual.Username)

		repo.AssertExpectations(t)
	})
	t.Run("Update error user not found", func(t *testing.T) {

		token := jwt.New(jwt.SigningMethodHS256)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "id user not found")
		assert.Empty(t, actual)
	})
	t.Run("Update error post not found", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(posting.Core{}, errors.New("not found")).Once()

		// Program service
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "Posting not found")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
	t.Run("Update error internal server", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(posting.Core{}, errors.New("internal server error")).Once()

		// Program service
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}
