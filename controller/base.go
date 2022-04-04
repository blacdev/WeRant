package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func User(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	return c.String(http.StatusOK, "name: "+email+" password: "+password)
}
