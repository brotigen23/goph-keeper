package model

import "time"

// Stores data about user accounts
type AccountData struct {
	ID         int
	UserID     int
	MetadataID int

	Login    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's text data
type TextData struct {
	ID         int
	UserID     int
	MetadataID int

	Data string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's binary data
type BinaryData struct {
	ID         int
	UserID     int
	MetadataID int

	Data []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Stores the user's bank card information
type CardData struct {
	ID         int
	UserID     int
	MetadataID int

	Number         string
	CardholderName string
	Expire         time.Time
	CVV            string

	CreatedAt time.Time
	UpdatedAt time.Time
}
