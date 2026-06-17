package main

import (
	"errors"
	"fmt"
)

func register(login, password string) error {
	if err := validateLogin(login); err != nil {
		fmt.Println("incorrect login")
		return err
	}

	if err := validatePassword(password); err != nil {
		fmt.Println("incorrect password")
		return err
	}

	if _, exist := users[login]; exist {
		return errors.New("user already is")
	}

	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	users[login] = User{
		Login:    login,
		Password: hash,
	}

	err = saveUserToFile(users[login])
	if err != nil {
		return err
	}
	return nil
}
