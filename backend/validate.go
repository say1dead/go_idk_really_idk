package main

import (
	"errors"
	"unicode"
)

func validateLogin(login string) error {
	if len(login) < 3 {
		return errors.New("login must be longer then 3 digits")
	}

	for _, ch := range login {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return errors.New("login must contain only digits and numbers")
		}
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be longer then 6 digits")
	}

	for _, ch := range password {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return errors.New("password must contain only digits and numbers")
		}
	}
	return nil
}
