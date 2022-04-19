package controller

import (
	"testing"
)

func TestUsernameGenratorLengthofCharacters(t *testing.T) {
	email := "owoborodeseye@gmail.com"
	password := "password"
	username := UserNameGenerator(email, password)
	actualLength := len(username)
	expectedLength := 28
	if actualLength != expectedLength {
		t.Errorf("Expected length of username to be %d, but got %d", expectedLength, actualLength)
	}

}

func TestUsernameGenratornotGeneratingEmptyCharacters(t *testing.T) {
	email := "qwertydgfhrjs@gmail.com"
	password := "password"
	username := UserNameGenerator(email, password)
	expectedOutput := ""

	if username == expectedOutput {
		t.Errorf("Expected output to be %s, but got %s", expectedOutput, username)
	}
}

func TestUsernameGenratornotgeneratingSameValueTwice(t *testing.T) {
	email := "qwertydgfhrjs@gmail.com"
	password := "password"
	username := UserNameGenerator(email, password)
	expectedOutput := UserNameGenerator(email, password)

	if username == expectedOutput {
		t.Errorf("Expected output to be %s, but got %s", expectedOutput, username)
	}
}
