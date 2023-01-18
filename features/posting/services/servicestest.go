package services

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// repo := mocks.NewPostingData(t)

	// srv := New(repo)

	// 	// Case: user melakukan pendaftaran akun baru
	// 	t.Run("Add successfully", func(t *testing.T) {
	// 		// Prgramming input and return repo

	// 		type SamplePosting struct {
	// 			ID        int
	// 			Postingan string
	// 			UserName  string
	// 			Image_url string
	// 			CreateAt  time.Time
	// 		}
	// 		sample := SamplePosting{
	// 			ID:        1,
	// 			Postingan: "reza",
	// 			Username:  "reza06",
	// 			Image_url: "12345",
	// 		}
	// 		input := PostUser.Core{
	// 			Postingan: sample.Postingan,
	// 			UserName:  sample.Username,
	// 			Image_url: sample.Image_url,
	// 			CreateAt:  time.Time,
	// 		}

	// 		// Program service

	// 		hashed, _ := helper.GeneratePassword(input.Password)
	// 		resData := user.Core{
	// 			Name:     input.Name,
	// 			Email:    input.Email,
	// 			Username: input.Username,
	// 			Password: hashed,
	// 		}
	// 		repo.On("Register", mock.Anything).Return(resData, nil).Once()
	// 		data, err := srv.Add(input)

	// 		assert.NoError(t, errCompare)
	// 		assert.Equal(t, data.ID, resData.ID)
	// 		repo.AssertExpectations(t)
	// 	})

	// 	t.Run("Masalah server", func(t *testing.T) {
	// 		type SampleUsers struct {
	// 			ID       int
	// 			Name     string
	// 			Email    string
	// 			Username string
	// 			Password string
	// 		}
	// 		sample := SampleUsers{
	// 			ID:       1,
	// 			Name:     "reza",
	// 			Email:    "reza@gmail.com",
	// 			Username: "reza06",
	// 			Password: "12345",
	// 		}
	// 		input := user.Core{
	// 			Name:     sample.Name,
	// 			Email:    sample.Email,
	// 			Username: sample.Username,
	// 			Password: sample.Password,
	// 		}

	// 		// Programming input and return repo
	// 		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()

	// 		// Program service
	// 		data, err := srv.Register(input)

	// 		// Test
	// 		assert.NotNil(t, err)
	// 		assert.EqualError(t, err, "terdapat masalah pada server")
	// 		assert.Empty(t, data)
	// 		repo.AssertExpectations(t)
	// 	})
}
