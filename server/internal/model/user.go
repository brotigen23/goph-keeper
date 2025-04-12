package model

import "time"

// The user's entity that stores the ID, login, and password
type User struct {
	ID int

	Login    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}
