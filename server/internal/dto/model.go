package dto

import (
	"time"
)

type BaseData struct {
	ID int `json:"id" example:"1"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
