package account

import "github.com/brotigen23/goph-keeper/server/internal/dto"

type PutRequest struct {
	ID int `json:"id" example:"1"`
	Model
}

type PutResponse struct {
	dto.BaseData
	Model
}
