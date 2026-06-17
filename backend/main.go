package main

import "fmt"

var users = map[string]User{}

func main() {
	err := loadUsersFromFile()

	if err != nil {
		fmt.Println("error with file: ", err)
		return
	}
	startWork()
}
