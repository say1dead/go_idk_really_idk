package main

func login(login, password string) (User, bool) {
	user, exist, err := getUserByLogin(login)
	if err != nil {
		return User{}, false
	}

	if !exist {
		return User{}, false
	}

	if !checkPassword(password, user.Password) {
		return User{}, false
	}

	return user, true
}
