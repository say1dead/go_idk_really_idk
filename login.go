package main

func login(login, password string) bool {
	user, exist := users[login]
	if !exist {
		return false
	}
	return checkPassword(password, user.Password)
}
