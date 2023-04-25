package main

import (
	"log"
	"taimekalender/back-end/driver"

	"github.com/gofrs/uuid"
)

func setSessionID(id int) string {
	sessionID, _ := uuid.NewV4()
	_, err := driver.DB.Exec("DELETE FROM sessions WHERE userId = ?", id)
	if err != nil {
		log.Println("Delete cookie error: ", err)
	}
	_, err = driver.DB.Exec("INSERT INTO sessions (uuid, userId) VALUES (?, ?) ", sessionID.String(), id)
	if err != nil {
		log.Println("Insert cookie error: ", err)
	}
	return sessionID.String()
}

func removeUUID(uuid string) {
	_, err := driver.DB.Exec("DELETE FROM sessions WHERE uuid = ?", uuid)
	if err != nil {
		log.Println("removeCookie error: ", err)
	}
}
