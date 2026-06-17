package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

var users = map[string]User{}

func main() {
	godotenv.Load()

	err := connectDB()
	if err != nil {
		fmt.Println("db connection error:", err)
		return
	}

	err = createUsersTable()
	if err != nil {
		fmt.Println("create table error:", err)
		return
	}

	startWork()
}
