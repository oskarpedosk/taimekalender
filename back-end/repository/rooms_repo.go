package repository

import (
	"taimekalender/back-end/driver"
	"taimekalender/back-end/models"
)

func GetRooms() ([]models.Room, error) {
	var rooms []models.Room
	rows, err := driver.DB.Query(`
		SELECT 
			*
		FROM 
			rooms
		ORDER BY
			"index" ASC;
		`)
	if err != nil {
		return rooms, err
	}

	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Index,
			&room.CreatedAt,
			&room.UpdatedAt)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func GetLastRoomIndex() (int, error) {
	var index int
	err := driver.DB.QueryRow(`SELECT COALESCE(MAX("index"), 0) FROM rooms`).Scan(&index)
	if err != nil {
		return index, err
	}

	return index, nil
}

func AddRoom(room models.Room) (models.Room, error) {
	index, err := GetLastRoomIndex()
	if err != nil {
		return room, err
	}

	result, err := driver.DB.Exec(`
		INSERT INTO
			rooms (name, "index")
		VALUES 
			(?, ?)`, room.Name, index+1)
	if err != nil {
		return room, err
	}

	roomID, err := result.LastInsertId()
	if err != nil {
		return room, err
	}

	room.ID = int(roomID)
	room.Index = index + 1

	return room, nil
}
