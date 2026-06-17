async function registerUser() {
	const firstName = document.getElementById("register-name").value
	const lastName = document.getElementById("register-surname").value
	const login = document.getElementById("register-login").value
	const password = document.getElementById("register-password").value

	const response = await fetch("/register", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded"
		},
		body: new URLSearchParams({
			firstName: firstName,
			lastName: lastName,
			login: login,
			password: password
		})
	})

	const text = await response.text()
	document.getElementById("message").textContent = text
}

async function loginUser() {
	const login = document.getElementById("login-login").value
	const password = document.getElementById("login-password").value

	const response = await fetch("/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded"
		},
		body: new URLSearchParams({
			login: login,
			password: password
		})
	})

	if (!response.ok) {
		const text = await response.text()
		document.getElementById("message").textContent = text
		return
	}

	const user = await response.json()

	document.getElementById("hello-text").textContent =
		"Привет " + user.firstName + " " + user.lastName

	hideAllPages()
	document.getElementById("account-page").style.display = "block"
}

function logout() {
	showChoice()
}

function hideAllPages() {
	document.getElementById("choice-page").style.display = "none"
	document.getElementById("register-page").style.display = "none"
	document.getElementById("login-page").style.display = "none"
	document.getElementById("account-page").style.display = "none"
}

function showChoice() {
	hideAllPages()
	document.getElementById("choice-page").style.display = "block"
	document.getElementById("message").textContent = ""
}

function showRegister() {
	hideAllPages()
	document.getElementById("register-page").style.display = "block"
	document.getElementById("message").textContent = ""
}

function showLogin() {
	hideAllPages()
	document.getElementById("login-page").style.display = "block"
	document.getElementById("message").textContent = ""
}