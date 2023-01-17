package handler

import "project/features/posting"

type PostingResponse struct {
	ID        uint   `json:"id"`
	Postingan string `json:"postingan"`
	UserName  string `json:"username"`
}
type AddPostingResponse struct {
	Postingan string `json:"postingan"`
}
type updatePostingResponse struct {
	Postingan string `json:"postingan"`
}

func ToResponse(feature string, posting posting.Core) interface{} {
	switch feature {
	case "add":
		return AddPostingResponse{
			Postingan: posting.Postingan,
		}
	case "update":
		return updatePostingResponse{
			Postingan: posting.Postingan,
		}
	default:
		return PostingResponse{
			ID:        posting.ID,
			Postingan: posting.Postingan,
			UserName:  posting.UserName,
		}
	}
}

func ListBookCoreToBookRespon(dataCore posting.Core) PostingResponse { // data user core yang ada di controller yang memanggil user repository
	return PostingResponse{
		ID:        dataCore.ID,
		Postingan: dataCore.Postingan,
		UserName:  dataCore.UserName,
	}
}
func ListBookCoreToBooksRespon(dataCore []posting.Core) []PostingResponse {
	var ResponData []PostingResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListBookCoreToBookRespon(value))
	}
	return ResponData
}
