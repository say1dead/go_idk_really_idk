package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
		<h1>Auth</h1>

		<h2>Register</h2>
		<form method="POST" action="/register">
			<input name="login" placeholder="login">
			<input name="password" placeholder="password" type="password">
			<button type="submit">Register</button>
		</form>

		<h2>Login</h2>
		<form method="POST" action="/login">
			<input name="login" placeholder="login">
			<input name="password" placeholder="password" type="password">
			<button type="submit">Login</button>
		</form>
	`)
}
