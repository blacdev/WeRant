package controller

import (
	"crypto/rand"

	"io"
	"net/http"
	"net/mail"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	Message string `json:"message"`
	Userid  string `json:"userid"`
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func UserRegistration(c echo.Context) error {
	email := c.FormValue("email")
	password, _ := HashPassword(c.FormValue("password"))

	if email == "" || password == "" {
		c.Response().Header().Set("Content-Type", "application/json")
		u := response{Message: "email and password are required"}
		return c.JSON(http.StatusBadRequest, u)

	}

	if !emailIsValid(email) {
		c.Response().Header().Set("Content-Type", "application/json")
		u := response{Message: "email is invalid"}
		return c.JSON(http.StatusBadRequest, u)
	}

	username := UserNameGenerator(email, password)

	c.Response().Header().Set("Content-Type", "application/json")
	u := response{Message: "user registered successfully", Userid: username}
	return c.JSON(http.StatusOK, u)
}

func UserLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	passwordMatch := CheckPasswordHash(password, "password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, response{Message: "username and password are required"})
	}

	if !passwordMatch {
		// return a json response
		return c.JSON(http.StatusBadRequest, response{Message: "password is invalid"})
	}

	return c.JSON(http.StatusOK, response{Message: "user logged in successfully", Userid: username})
}

func UserNameGenerator(email, password string) (username string) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	username = "Ranter-"
	email = email[:strings.Index(email, "@")]
	userDetailsEncoder := len(email) + len(password)
	b := make([]byte, userDetailsEncoder)
	n, err := io.ReadAtLeast(rand.Reader, b, userDetailsEncoder)
	if n != userDetailsEncoder {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	username = username + string(b)
	return username
}

func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
