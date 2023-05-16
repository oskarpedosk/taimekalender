package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Room struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Index     int       `json:"index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Plant struct {
	ID                  int       `json:"id"`
	RoomID              int       `json:"room_id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Index               int       `json:"index"`
	WateringInterval    int       `json:"watering_interval"`
	Watered             time.Time `json:"watered"`
	FertilizingInterval int       `json:"fertilizing_interval"`
	Fertilized          time.Time `json:"fertilized"`
	Transplanting       string    `json:"transplanting"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
