package handler

import "project/features/posting"

type AddUpdatePostingRequest struct {
	Postingan string `form:"postingan"`
}

func ToCore(data interface{}) *posting.Core {
	res := posting.Core{}

	switch data.(type) {
	case AddUpdatePostingRequest:
		cnv := data.(AddUpdatePostingRequest)
		res.Postingan = cnv.Postingan

	default:
		return nil
	}

	return &res
}
