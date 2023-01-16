package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Name     string
	Email    string
	Username string
	Password string
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Deactive() echo.HandlerFunc
	Update() echo.HandlerFunc
	AllUser() echo.HandlerFunc
}

type UserService interface {
	Login(email, password string) (string, Core, error)
	Register(newUser Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core) (Core, error)
	Deactive(token interface{}) (Core, error)
	AllUser() ([]Core, error)
}

type UserData interface {
	Login(email string) (Core, error)
	Register(newUser Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	Deactive(id uint) (Core, error)
	AllUser() ([]Core, error)
}
