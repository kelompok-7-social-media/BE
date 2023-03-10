package handler

import (
	"project/features/komentar"
)

type KomentarResponse struct {
	ID        uint   `json:"id"`
	Comment   string `json:"comment"`
	PostingID uint   `json:"posting_id"`
	CreatedAt string `json:"tanggal"`
}
type KomentarsResponse struct {
	ID        uint   `json:"id"`
	Comment   string `json:"comment"`
	PostingID uint   `json:"posting_id"`
	Username  string `json:"username"`
}
type AddKomentarResponse struct {
	Comment   string `json:"comment"`
	PostingID uint   `json:"posting_id"`
	CreatedAt string `json:"tanggal"`
}

type updatePostingResponse struct {
	Comment   string `json:"comment"`
	PostingID uint   `json:"posting_id"`
}

func ToResponse(feature string, komentar komentar.Core) interface{} {
	switch feature {
	case "add":
		return AddKomentarResponse{
			Comment:   komentar.Comment,
			CreatedAt: komentar.CreatedAt,
			PostingID: komentar.PostingID,
		}
	case "update":
		return updatePostingResponse{
			Comment: komentar.Comment,

			PostingID: komentar.PostingID,
		}
	default:
		return KomentarResponse{
			ID:        komentar.ID,
			Comment:   komentar.Comment,
			CreatedAt: komentar.CreatedAt,
			PostingID: komentar.PostingID,
		}
	}
}

func ListCommentCoreToCommentRespon(dataCore komentar.Core) KomentarsResponse { // data user core yang ada di controller yang memanggil user repository
	return KomentarsResponse{
		ID:        dataCore.ID,
		Comment:   dataCore.Comment,
		PostingID: dataCore.PostingID,
		Username:  dataCore.Username,
	}
}
func ListCommentCoreToCommentsRespon(dataCore []komentar.Core) []KomentarsResponse {
	var ResponData []KomentarsResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListCommentCoreToCommentRespon(value))
	}
	return ResponData
}
