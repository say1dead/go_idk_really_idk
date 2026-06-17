package main

import (
	"fmt"
	"net/http"
)

func startWork() {
	err := loadUsersFromFile()
	if err != nil {
		fmt.Println("Ошибка загрузки пользователей:", err)
		return
	}

	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
