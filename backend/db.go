package main

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func connectDB() error {
	databaseUrl := os.Getenv("DATABASE_URL")

	var err error
	db, err = sql.Open("pgx", databaseUrl)

	if err != nil {
		return err
	}

	return db.Ping()
}

func createUsersTable() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		surname TEXT NOT NULL,
		login TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now()
	);
	`)
	return err
}

func saveUserToDb(user User) error {
	_, err := db.Exec(`
	INSERT INTO users (name, surname, login, password)
	VALUES ($1, $2, $3, $4)
	`,
		user.Name,
		user.Surname,
		user.Login,
		user.Password,
	)

	return err
}

func getUserByLogin(login string) (User, bool, error) {
	var user User

	err := db.QueryRow(`
		SELECT name, surname, login, password
		FROM users
		WHERE login = $1
	`, login).Scan(
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
	)

	if err == sql.ErrNoRows {
		return User{}, false, nil
	}

	if err != nil {
		return User{}, false, err
	}

	return user, true, nil
}
