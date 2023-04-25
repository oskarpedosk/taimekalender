package repository

import (
	"database/sql"
	"log"
	"taimekalender/back-end/driver"

	"golang.org/x/crypto/bcrypt"
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
	err := driver.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&dbPassword)
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
