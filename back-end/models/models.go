package models

import "time"

type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"passwordConfirm"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	DateOfBirth     time.Time `json:"dateOfBirth"`
	Image           string    `json:"image"`
	Nickname        string    `json:"nickname"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UUID            string    `json:"uuid"`
}

type Room struct {
}

type Plant struct {
}
