package repository

import (
	"log"
	"taimekalender/back-end/driver"
	"taimekalender/back-end/models"
)

func GetPlants() ([]models.Plant, error) {
	var plants []models.Plant
	rows, err := driver.DB.Query(`
		SELECT 
			*
		FROM 
			plants
		ORDER BY
			"index" ASC
		`)
	if err != nil {
		return plants, err
	}

	defer rows.Close()
	for rows.Next() {
		var plant models.Plant
		err := rows.Scan(
			&plant.ID,
			&plant.RoomID,
			&plant.Name,
			&plant.Description,
			&plant.Index,
			&plant.WateringInterval,
			&plant.Watered,
			&plant.FertilizingInterval,
			&plant.Fertilized,
			&plant.Transplanting,
			&plant.CreatedAt,
			&plant.UpdatedAt)
		if err != nil {
			return plants, err
		}
		plants = append(plants, plant)
	}

	return plants, nil
}

func GetRoomPlants(roomID int) ([]models.Plant, error) {
	var plants []models.Plant
	rows, err := driver.DB.Query(`
		SELECT 
			*
		FROM 
			plants
		WHERE
			room_id = ?
		ORDER BY
			"index" ASC
		`, roomID)
	if err != nil {
		return plants, err
	}

	defer rows.Close()
	for rows.Next() {
		var plant models.Plant
		err := rows.Scan(
			&plant.ID,
			&plant.RoomID,
			&plant.Name,
			&plant.Description,
			&plant.Index,
			&plant.WateringInterval,
			&plant.Watered,
			&plant.FertilizingInterval,
			&plant.Fertilized,
			&plant.Transplanting,
			&plant.CreatedAt,
			&plant.UpdatedAt)
		if err != nil {
			return plants, err
		}
		plants = append(plants, plant)
	}

	return plants, nil
}

func GetLastPlantIndex() (int, error) {
	var index int
	err := driver.DB.QueryRow(`SELECT COALESCE(MAX("index"), 0) FROM plants`).Scan(&index)
	if err != nil {
		return index, err
	}

	return index, nil
}

func AddPlant(plant models.Plant) (models.Plant, error) {
	log.Println(plant)
	index, err := GetLastPlantIndex()
	if err != nil {
		return plant, err
	}

	result, err := driver.DB.Exec(`
		INSERT INTO
			plants (room_id, name, description, "index", watering_interval, watered, fertilizing_interval, fertilized, transplanting)
		VALUES 
			(?, ?, ?, ?, ?, ?, ?, ?, ?)`, plant.RoomID, plant.Name, plant.Description, index+1, plant.WateringInterval, plant.Watered, plant.FertilizingInterval, plant.Fertilized, plant.Transplanting)
	if err != nil {
		return plant, err
	}

	plantID, err := result.LastInsertId()
	if err != nil {
		return plant, err
	}

	plant.ID = int(plantID)
	plant.Index = index + 1

	return plant, nil
}
