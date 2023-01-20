package handler

import (
	"project/features/posting"
	"time"
)

type PostingResponse struct {
	ID        uint   `json:"id"`
	Postingan string `json:"postingan"`
	UserName  string `json:"username"`
	Image_url string `json:"image_url"`
}
type AddPostingResponse struct {
	Postingan string    `json:"postingan"`
	Image_url string    `json:"image_url"`
	CreatedAt time.Time `json:"tanggal"`
}
type updatePostingResponse struct {
	Postingan string `json:"postingan"`
	Image_url string `json:"image_url"`
}

func ToResponse(feature string, posting posting.Core) interface{} {
	switch feature {
	case "add":
		return AddPostingResponse{
			Postingan: posting.Postingan,
			CreatedAt: posting.CreatedAt,
			Image_url: posting.Image_url,
		}
	case "update":
		return updatePostingResponse{
			Postingan: posting.Postingan,
			Image_url: posting.Image_url,
		}
	default:
		return PostingResponse{
			ID:        posting.ID,
			Postingan: posting.Postingan,
			UserName:  posting.Username,
		}
	}
}

func ListPostCoreToPostRespon(dataCore posting.Core) PostingResponse { // data user core yang ada di controller yang memanggil user repository
	return PostingResponse{
		ID:        dataCore.ID,
		Postingan: dataCore.Postingan,
		Image_url: dataCore.Image_url,
		UserName:  dataCore.Username,
	}
}
func ListPostCoreToPostsRespon(dataCore []posting.Core) []PostingResponse {
	var ResponData []PostingResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListPostCoreToPostRespon(value))
	}
	return ResponData
}
