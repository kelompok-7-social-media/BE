package handler

import "project/features/posting"

type AddUpdatePostingRequest struct {
	Postingan string `form:"postingan"`
	Image_url string `form:"image"`
}

func (data *AddUpdatePostingRequest) reqToCore() posting.Core {
	return posting.Core{
		Postingan: data.Postingan,
		Image_url: data.Image_url,
	}
}

func ToCore(data interface{}) *posting.Core {
	res := posting.Core{}

	switch data.(type) {
	case AddUpdatePostingRequest:
		cnv := data.(AddUpdatePostingRequest)
		res.Postingan = cnv.Postingan
		res.Image_url = cnv.Image_url

	default:
		return nil
	}

	return &res
}
