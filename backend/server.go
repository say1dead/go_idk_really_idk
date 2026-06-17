package main

import (
	"fmt"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method incorrect", http.StatusMethodNotAllowed)
		return
	}

	loginInput := r.FormValue("login")
	passwordInput := r.FormValue("password")

	err := register(loginInput, passwordInput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "user created")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method incorrect", http.StatusMethodNotAllowed)
		return
	}

	loginInput := r.FormValue("login")
	passwordInput := r.FormValue("password")

	if login(loginInput, passwordInput) {
		fmt.Fprintln(w, "login succes")
	} else {
		http.Error(w, "wrong login or password", http.StatusUnauthorized)
	}
}
