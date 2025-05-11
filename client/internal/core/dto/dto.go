package dto

import (
	"time"
)

type BaseDTO struct {
	ID int `json:"id" example:"1"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseError struct {
	Msg string `json:"msg"`
} //@name DTO.ResponseError
