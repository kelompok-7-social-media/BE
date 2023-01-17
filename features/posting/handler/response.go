package handler

import "project/features/posting"

type PostingResponse struct {
	ID        uint   `json:"id"`
	Postingan string `json:"postingan"`
	UserName  string `json:"username"`
	Image     string `json:"image"`
}
type AddPostingResponse struct {
	Postingan string `json:"postingan"`
	Image     string `json:"image"`
}
type updatePostingResponse struct {
	Postingan string `json:"postingan"`
	Image     string `json:"image"`
}

func ToResponse(feature string, posting posting.Core) interface{} {
	switch feature {
	case "add":
		return AddPostingResponse{
			Postingan: posting.Postingan,
			Image:     posting.Image,
		}
	case "update":
		return updatePostingResponse{
			Postingan: posting.Postingan,
			Image:     posting.Image,
		}
	default:
		return PostingResponse{
			ID:        posting.ID,
			Postingan: posting.Postingan,
			Image:     posting.Image,
			UserName:  posting.UserName,
		}
	}
}

func ListBookCoreToBookRespon(dataCore posting.Core) PostingResponse { // data user core yang ada di controller yang memanggil user repository
	return PostingResponse{
		ID:        dataCore.ID,
		Postingan: dataCore.Postingan,
		Image:     dataCore.Image,
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
