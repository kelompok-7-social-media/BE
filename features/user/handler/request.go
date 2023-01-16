package handler

import "api/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"nama" form:"nama"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name     string `json:"nama" form:"nama"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Username= cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
