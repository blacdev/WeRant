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

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func UserRegistration(c echo.Context) error {
	email := c.FormValue("email")
	password, _ := HashPassword(c.FormValue("password"))

	if email == "" || password == "" {
		return c.String(http.StatusBadRequest, "email and password are required")
	}

	if !emailIsValid(email) {
		return c.String(http.StatusBadRequest, "email is invalid")
	}

	username := userNameGenerator(email, password)

	return c.String(http.StatusOK, "username: "+username)
}

func UserLogin(c echo.Context) error {
	username := c.FormValue("email")
	password := c.FormValue("password")

	passwordMatch := CheckPasswordHash(password, "password")

	if username == "" || password == "" {
		return c.String(http.StatusBadRequest, "email and password are required")
	}

	if !passwordMatch {
		return c.String(http.StatusBadRequest, "password is invalid")
	}

	return c.String(http.StatusOK, "username: "+username)
}

func userNameGenerator(email, password string) (username string) {
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
