package handler

import (
	"project/features/komentar"
)

type AddUpdateKomenRequest struct {
	Pesan  string `form:"pesan"`
	PostID uint   `form:"postid"`
}
type UpdateKomenRequest struct {
	Pesan  string `form:"pesan"`
	PostID uint   `form:"postid"`
}

func (data *AddUpdateKomenRequest) reqToCore() komentar.Core {
	return komentar.Core{
		Pesan:     data.Pesan,
		PostingID: data.PostID,
	}
}

func ToCore(data interface{}) *komentar.Core {
	res := komentar.Core{}

	switch data.(type) {
	case AddUpdateKomenRequest:
		cnv := data.(AddUpdateKomenRequest)
		res.Pesan = cnv.Pesan
		res.PostingID = cnv.PostID

	default:
		return nil
	}

	return &res
}
