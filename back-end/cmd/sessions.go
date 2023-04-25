package main

import (
	"log"

	"github.com/gofrs/uuid"
)

func setSessionID(id int) string {
	sessionID, _ := uuid.NewV4()
	_, err := database.Exec("DELETE FROM sessions WHERE userId = ?", id)
	if err != nil {
		log.Println("Delete cookie error: ", err)
	}
	_, err = database.Exec("INSERT INTO sessions (uuid, userId) VALUES (?, ?) ", sessionID.String(), id)
	if err != nil {
		log.Println("Insert cookie error: ", err)
	}
	return sessionID.String()
}

func removeUUID(uuid string) {
	_, err := database.Exec("DELETE FROM sessions WHERE uuid = ?", uuid)
	if err != nil {
		log.Println("removeCookie error: ", err)
	}
}
