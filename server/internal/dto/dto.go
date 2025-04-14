package dto

import "time"

type Validator interface {
	Validate() error
}

type BaseData struct {
	ID int `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
