package dto

import "time"

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BaseData struct {
	ID int `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Metadata struct {
	BaseData
	Metadata string
}

type AccountsGet struct {
	BaseData
	AccountData
	Metadata Metadata
}
