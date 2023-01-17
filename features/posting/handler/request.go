package handler

import "project/features/posting"

type AddUpdatePostingRequest struct {
	Postingan string `form:"postingan`
	Image     string `form:"image"`
}

func ToCore(data interface{}) *posting.Core {
	res := posting.Core{}

	switch data.(type) {
	case AddUpdatePostingRequest:
		cnv := data.(AddUpdatePostingRequest)
		res.Postingan = cnv.Postingan
		res.Image = cnv.Image

	default:
		return nil
	}

	return &res
}
