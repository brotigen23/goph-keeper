package domain

import "time"

type AccountData struct {
	ID         int
	UserID     int
	MetadataID int

	Login    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}
