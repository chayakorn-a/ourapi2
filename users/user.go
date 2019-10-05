package users

import (
	"net/http"

	"github.com/chayakorn-a/ourapi2/typicode"
	"github.com/labstack/echo"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Decoder interface {
	Decode(result interface{}) error
}

type userAPI struct {
	service Decoder
}

func (up *userAPI) GetUsers(c echo.Context) error {
	uu := []User{}

	err := up.service.Decode(&uu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, uu)
}

func GetUsers(c echo.Context) error {
	api := &userAPI{
		//service: &typicode{},
		service: typicode.Get("/users"),
	}

	return api.GetUsers(c)
}
