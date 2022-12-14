package domain

import (
	user "project3/group3/feature/users"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Phone     string
	Role      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//logic
type UserUseCase interface {
	AddUser(newUser User) (row int, err error)
	Login(auth user.LoginModel) (data map[string]interface{}, err error)
	GetProfile(id int) (User, error)
	DeleteCase(userID int) (row int, err error)
	UpdateCase(input User, idUser int) (row int, err error)
}

//query
type UserData interface {
	Insert(newUser User) (row int, err error)
	LoginData(authData user.LoginModel) (data map[string]interface{}, err error)
	GetSpecific(userID int) (User, error)
	DeleteData(userID int) (row int, err error)
	UpdateData(data map[string]interface{}, idUser int) (row int, err error)
}

//handler
type UserHandler interface {
	InsertUser() echo.HandlerFunc
	LoginAuth() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	DeleteById() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}
