package handlers

import (
	"fmt"
	"net/http"
	"taimekalender/back-end/helpers"
	"taimekalender/back-end/models"
	"taimekalender/back-end/repository"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := repository.GetRooms()
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cannot get rooms %v", err), http.StatusBadRequest)
		return
	}

	_ = helpers.WriteJson(w, http.StatusOK, rooms, nil)
}

func AddRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	_ = helpers.ReadJson(w, r, &room)

	room, err := repository.AddRoom(room)
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cannot add room %v", err), http.StatusBadRequest)
		return
	}

	_ = helpers.WriteJson(w, http.StatusOK, room, nil)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {

}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {

}
