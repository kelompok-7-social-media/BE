package handler

import (
	comment "project/features/komentar"
)

type CommentResponse struct {
	ID         uint   `json:"id" form:"id"`
	Posting_ID uint   `json:"posting"`
	Comment    string `json:"comment"`
}
type AddCommentResponse struct {
	Comment string `json:"comment"`
}
type updateCommentResponse struct {
	Comment string `json:"comment"`
}

func ToResponse(feature string, comment comment.Core) interface{} {
	switch feature {
	case "add":
		return AddCommentResponse{
			Comment: comment.Comment,
		}
	case "update":
		return updateCommentResponse{
			Comment: comment.Comment,
		}
	default:
		return CommentResponse{
			ID:         comment.ID,
			Posting_ID: comment.Posting_ID,
			Comment:    comment.Comment,
		}
	}
}

func CoreToRespon(dataCore comment.Core) CommentResponse {
	return CommentResponse{
		ID:         dataCore.ID,
		Posting_ID: dataCore.Posting_ID,
		Comment:    dataCore.Comment,
	}
}
func CoreToResponses(dataCore []comment.Core) []CommentResponse {
	var ResponData []CommentResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, CoreToRespon(value))
	}
	return ResponData
}
