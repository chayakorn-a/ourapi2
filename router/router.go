package router

import (
	"github.com/chayakorn-a/ourapi2/users"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/users", users.GetUsers)
	return e
}
