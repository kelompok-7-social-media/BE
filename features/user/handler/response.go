package handler

import (
	"api/features/user"
)

type UserReponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"nama"`
	HP    string `json:"hp"`
	Email string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:    data.ID,
		Name:  data.Name,
		HP:    data.HP,
		Email: data.Email,
	}
}

// func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
// 	resp := map[string]interface{}{}
// 	// if len(data) < 2 {
// 	// 	resp["data"] = ToResponse(data[0].(user.Core))
// 	// } else {
// 	// 	resp["data"] = ToResponse(data[0].(user.Core))
// 	// 	resp["token"] = data[1].(string)
// 	// }
// 	switch len(data) {
// 	case 1:
// 		resp["data"] = data[0]
// 	case 2:
// 		resp["token"] = data[1].(string)
// 		resp["data"] = data[0]
// 	}
// 	if message != "" {
// 		resp["message"] = message
// 	}

// 	return code, resp
// }

// func PrintErrorResponse(msg string) (int, interface{}) {
// 	resp := map[string]interface{}{}
// 	code := -1
// 	if msg != "" {
// 		resp["message"] = msg
// 	}

// 	if strings.Contains(msg, "server") {
// 		code = http.StatusInternalServerError
// 	} else if strings.Contains(msg, "format") {
// 		code = http.StatusBadRequest
// 	} else if strings.Contains(msg, "not found") {
// 		code = http.StatusNotFound
// 	}

// 	return code, resp
// }
