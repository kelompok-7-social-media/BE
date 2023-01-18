package handler

import "project/features/user"

type UserReponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	}
}

func ToResponses(data user.Core) UserReponse {
	return UserReponse{

		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
	}
}
func fromCoreList(dataCore []user.Core) []UserReponse {
	var dataResponse []UserReponse

	for _, v := range dataCore {
		dataResponse = append(dataResponse, ToResponse(v))
	}
	return dataResponse
}
