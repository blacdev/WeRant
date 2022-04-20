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

// response is used to return a json response
type response struct {
	Message string `json:"message"`
	Userid  string `json:"userid"`
}

// Health controller is used to check the health of the application
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

// user registration controller registers the user and returns a json response
func UserRegistration(c echo.Context) error {
	email := c.FormValue("email")

	// password hash
	password, _ := HashPassword(c.FormValue("password"))

	// check if email or password is empty
	if email == "" || password == "" {
		c.Response().Header().Set("Content-Type", "application/json")
		u := response{Message: "email and password are required"}
		return c.JSON(http.StatusBadRequest, u)

	}

	// check if email is valid
	if !emailIsValid(email) {
		c.Response().Header().Set("Content-Type", "application/json")
		u := response{Message: "email is invalid"}
		return c.JSON(http.StatusBadRequest, u)
	}

	// generate username from email and password
	username := UserNameGenerator(email, password)

	c.Response().Header().Set("Content-Type", "application/json")
	u := response{Message: "user registered successfully", Userid: username}
	return c.JSON(http.StatusOK, u)
}

// UserLogin controller is used to login the user and returns a json response
func UserLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// hash user input password
	passwordMatch := CheckPasswordHash(password, "password")

	// check if username or password is empty
	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, response{Message: "username and password are required"})
	}

	// check if password hash matches the password
	if !passwordMatch {
		// return a json response
		return c.JSON(http.StatusBadRequest, response{Message: "password is invalid"})
	}

	return c.JSON(http.StatusOK, response{Message: "user logged in successfully", Userid: username})
}

// UserNameGenerator generates a username from email and password
func UserNameGenerator(email, password string) (username string) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	username = "Ranter-"

	// reads email name till the @ sign
	email = email[:strings.Index(email, "@")]

	// get length of email and password
	userDetailsEncoder := len(email) + len(password)

	// make a byte array of the length of email and password
	b := make([]byte, userDetailsEncoder)

	// read at least the length of email and password
	n, err := io.ReadAtLeast(rand.Reader, b, userDetailsEncoder)
	if n != userDetailsEncoder {
		panic(err)
	}

	// loop through the byte array and assign a random number to each byte
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	// return the username
	username = username + string(b)
	return username
}

// emailIsValid checks if the email is valid
func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if the password matches the hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
