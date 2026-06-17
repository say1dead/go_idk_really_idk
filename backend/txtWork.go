package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func saveUserToFile(user User) error {
	file, err := os.OpenFile("users.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%s:%s\n", user.Login, user.Password)
	return err
}

func loadUsersFromFile() error {
	file, err := os.Open("users.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			continue
		}

		login := parts[0]
		password := parts[1]

		users[login] = User{
			Login:    login,
			Password: password,
		}
	}

	return scanner.Err()
}
