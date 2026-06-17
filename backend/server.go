package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method incorrect", http.StatusMethodNotAllowed)
		return
	}

	nameInput := r.FormValue("firstName")
	surnameInput := r.FormValue("lastName")
	loginInput := r.FormValue("login")
	passwordInput := r.FormValue("password")

	err := register(nameInput, surnameInput, loginInput, passwordInput)

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
	user, ok := login(loginInput, passwordInput)

	if ok {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"firstName": user.Name,
			"lastName":  user.Surname,
		})

		return
	} else {
		http.Error(w, "wrong login or password", http.StatusUnauthorized)
	}
}
