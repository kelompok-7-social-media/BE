package services

// import (
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestAddComment(t *testing.T) {
// 	repo := mocks.NewRepoInterface(t)
// 	t.Run("success add comments", func(t *testing.T) {
// 		repo.On("GetSingleStatus", mock.Anything).Return(nil, nil).Once()
// 		repo.On("AddComment", mock.Anything).Return(mentee.CommentsCore{}, nil).Once()
// 		srv := New(repo)
// 		input := mentee.CommentsCore{IdStatus: 1, Caption: "Aku bingung ini kenapa"}
// 		res, err := srv.Insert(input)
// 		assert.Empty(t, res)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("failed add status", func(t *testing.T) {
// 		repo.On("GetSingleStatus", mock.Anything).Return(nil, nil).Once()
// 		repo.On("AddComment", mock.Anything).Return(mentee.CommentsCore{}, errors.New("Error")).Once()
// 		input := mentee.CommentsCore{
// 			IdStatus: 1,
// 			Caption:  "asasas",
// 		}
// 		srv := New(repo)
// 		_, err := srv.Insert(input)
// 		assert.Error(t, err)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("failed add status", func(t *testing.T) {
// 		repo.On("GetSingleStatus", mock.Anything).Return(errors.New("failed get status"), errors.New("failed get status")).Once()
// 		input := mentee.CommentsCore{

// 			Caption: "asasas",
// 		}
// 		srv := New(repo)
// 		res, err := srv.Insert(input)
// 		assert.Error(t, err)
// 		assert.NotNil(t, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Failed length not valid", func(t *testing.T) {
// 		srv := New(repo)
// 		input := mentee.CommentsCore{
// 			IdStatus: 1,
// 			Caption:  "as",
// 		}
// 		res, err := srv.Insert(input)
// 		assert.Empty(t, res)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "failed add your comment check charancter len")
// 		repo.AssertExpectations(t)
// 	})

// }
