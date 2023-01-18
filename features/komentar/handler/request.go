package handler

import (
	"project/features/komentar"
)

type AddKomenRequest struct {
	Comment string `form:"comment"`
	PostID  uint   `form:"postid"`
}
type UpdateKomenRequest struct {
	Pesan  string `form:"pesan"`
	PostID uint   `form:"postid"`
}

func (data *AddKomenRequest) reqToCore() komentar.Core {
	return komentar.Core{
		Comment:   data.Comment,
		PostingID: data.PostID,
	}
}

func ToCore(data interface{}) *komentar.Core {
	res := komentar.Core{}

	switch data.(type) {
	case AddKomenRequest:
		cnv := data.(AddKomenRequest)
		res.Comment = cnv.Comment
		res.PostingID = cnv.PostID

	default:
		return nil
	}

	return &res
}
