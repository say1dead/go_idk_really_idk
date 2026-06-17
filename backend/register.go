package main

import (
	"errors"
	"fmt"
)

func register(name, surname, login, password string) error {
	if err := validateLogin(login); err != nil {
		fmt.Println("incorrect login")
		return err
	}

	if err := validatePassword(password); err != nil {
		fmt.Println("incorrect password")
		return err
	}

	if err := validateName(name); err != nil {
		fmt.Println("incorrect name")
		return err
	}

	if err := validateSurname(surname); err != nil {
		fmt.Println("incorrect surname")
		return err
	}

	_, exist, err := getUserByLogin(login)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("user already is")
	}

	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := User{
		Name:     name,
		Surname:  surname,
		Login:    login,
		Password: hash,
	}

	return saveUserToDb(user)
}
