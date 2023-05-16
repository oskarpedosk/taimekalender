package repository

import (
	"taimekalender/back-end/driver"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateViaPassword(username, password string) (int, bool, error) {
	var userID int
	var dbPassword string
	err := driver.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &dbPassword)
	if err != nil {
		return 0, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return 0, false, err
	}

	return userID, true, nil
}

func AuthenticateViaCookie(uuid string) (bool, error) {
	var userID int
	err := driver.DB.QueryRow("SELECT id FROM sessions where uuid = ?", uuid).Scan(&userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
