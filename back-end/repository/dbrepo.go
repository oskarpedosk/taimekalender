package repository

import (
	"database/sql"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(email, password string) bool {
	var dbPassword string
	err := database.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Println("login() error: ", err)
	}
	if CheckPasswordHash(password, dbPassword) {
		return true
	}
	return false
}
