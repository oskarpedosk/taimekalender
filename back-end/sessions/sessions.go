package sessions

import (
	"taimekalender/back-end/driver"

	"github.com/gofrs/uuid"
)

func Set(userID int) (string, error) {
	sessionID, _ := uuid.NewV4()
	_, err := driver.DB.Exec("DELETE FROM sessions WHERE userId = ?", userID)
	if err != nil {
		return "", err
	}

	_, err = driver.DB.Exec("INSERT INTO sessions (uuid, userId) VALUES (?, ?) ", sessionID.String(), userID)
	if err != nil {
		return "", err
	}

	return sessionID.String(), nil
}

func Remove(uuid string) error {
	_, err := driver.DB.Exec("DELETE FROM sessions WHERE uuid = ?", uuid)
	if err != nil {
		return err
	}

	return nil
}
