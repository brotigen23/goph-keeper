package dto

import "time"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Account struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`

	Login    string `json:"login"`
	Password string `json:"password"`

	MetadataID int    `json:"metadata_id"`
	Metadata   string `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
