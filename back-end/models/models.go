package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}

type Room struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Index string `json:"index"`
}

type Plant struct {
	ID                  int       `json:"id"`
	RoomID              int       `json:"room_id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Index               string    `json:"index"`
	Watered             time.Time `json:"watered"`
	WateringInterval    int       `json:"watering_interval"`
	Fertilized          time.Time `json:"fertilized"`
	FertilizingInterval int       `json:"fertilizing_interval"`
	Transplanting       string    `json:"transplanting"`
}
