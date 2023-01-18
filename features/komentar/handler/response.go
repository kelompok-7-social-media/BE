package handler

import (
	"project/features/komentar"
	"time"
)

type KomentarResponse struct {
	ID        uint      `json:"id"`
	Pesan     string    `json:"pesan"`
	PostingID uint      `json:"posting_id"`
	CreatedAt time.Time `json:"tanggal"`
}
type AddKomentarResponse struct {
	Pesan     string    `json:"pesan"`
	PostingID uint      `json:"posting_id"`
	CreatedAt time.Time `json:"tanggal"`
}

type updatePostingResponse struct {
	Pesan     string    `json:"pesan"`
	PostingID uint      `json:"posting_id"`
	CreatedAt time.Time `json:"tanggal"`
}

func ToResponse(feature string, komentar komentar.Core) interface{} {
	switch feature {
	case "add":
		return AddKomentarResponse{
			Pesan:     komentar.Pesan,
			CreatedAt: komentar.CreatedAt,
			PostingID: komentar.PostingID,
		}
	case "update":
		return updatePostingResponse{
			Pesan:     komentar.Pesan,
			CreatedAt: komentar.CreatedAt,
			PostingID: komentar.PostingID,
		}
	default:
		return KomentarResponse{
			ID:        komentar.ID,
			Pesan:     komentar.Pesan,
			CreatedAt: komentar.CreatedAt,
			PostingID: komentar.PostingID,
		}
	}
}

// func ListBookCoreToBookRespon(dataCore posting.Core) PostingResponse { // data user core yang ada di controller yang memanggil user repository
// 	return PostingResponse{
// 		ID:        dataCore.ID,
// 		Postingan: dataCore.Postingan,
// 		UserName:  dataCore.UserName,
// 	}
// }
// func ListBookCoreToBooksRespon(dataCore []posting.Core) []PostingResponse {
// 	var ResponData []PostingResponse

// 	for _, value := range dataCore {
// 		ResponData = append(ResponData, ListBookCoreToBookRespon(value))
// 	}
// 	return ResponData
// }
