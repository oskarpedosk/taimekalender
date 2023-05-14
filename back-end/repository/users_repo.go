package repository

import (
	"taimekalender/back-end/driver"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username, password string) (int, bool, error) {
	var userID int
	var dbPassword string
	err := driver.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &dbPassword)
	if err != nil {
		return 0, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
	if err != nil {
		return 0, false, err
	}

	return userID, true, nil
}
