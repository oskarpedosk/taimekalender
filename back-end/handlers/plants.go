package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"taimekalender/back-end/helpers"
	"taimekalender/back-end/models"
	"taimekalender/back-end/repository"

	"github.com/go-chi/chi/v5"
)

func GetPlants(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "roomID"))
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("invalid room id %v", err), http.StatusBadRequest)
		return
	}

	var plants []models.Plant

	if roomID == 0 {
		plants, err = repository.GetPlants()
		if err != nil {
			_ = helpers.ErrorJson(w, fmt.Errorf("cannot get all plant %v", err), http.StatusBadRequest)
			return
		}
	} else {
		plants, err = repository.GetRoomPlants(roomID)
		if err != nil {
			_ = helpers.ErrorJson(w, fmt.Errorf("cannot get room plants %v", err), http.StatusBadRequest)
			return
		}
	}

	_ = helpers.WriteJson(w, http.StatusOK, plants, nil)
}

func AddPlant(w http.ResponseWriter, r *http.Request) {
	var plant models.Plant
	_ = helpers.ReadJson(w, r, &plant)

	plant, err := repository.AddPlant(plant)
	if err != nil {
		_ = helpers.ErrorJson(w, fmt.Errorf("cannot add plant %v", err), http.StatusBadRequest)
		return
	}

	_ = helpers.WriteJson(w, http.StatusOK, plant, nil)
}

func UpdatePlant(w http.ResponseWriter, r *http.Request) {

}

func DeletePlant(w http.ResponseWriter, r *http.Request) {

}
