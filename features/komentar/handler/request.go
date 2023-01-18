package handler

import comment "project/features/komentar"

type AddCommentRequest struct {
	Comment string `form:"comment"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	switch data.(type) {
	case AddCommentRequest:
		cnv := data.(AddCommentRequest)
		res.Comment = cnv.Comment

	default:
		return nil
	}

	return &res
}
