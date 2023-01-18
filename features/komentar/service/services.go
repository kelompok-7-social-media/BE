package services

import (
	"errors"
	comment "project/features/komentar"
	"project/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type commentSrv struct {
	data     comment.CommentData
	validasi *validator.Validate
}

func New(cd comment.CommentData) comment.CommentService {
	return &commentSrv{
		data:     cd,
		validasi: validator.New(),
	}
}
func (cs *commentSrv) Add(token interface{}, newComment comment.Core) (comment.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comment.Core{}, errors.New("user not found")
	}
	res, err := cs.data.Add(userID, newComment)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Comment not found"
		} else {
			msg = "internal server error"
		}
		return comment.Core{}, errors.New(msg)
	}

	return res, nil
}
