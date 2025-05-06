package binarydto

import "github.com/brotigen23/goph-keeper/server/internal/dto"

// ***************************
// * Model
// ***************************
type DTO struct {
	Data []byte `json:"data" example:"data"`

	Metadata string `json:"metadata" example:"metadata"`
}

// ***************************
// * POST
// ***************************
type PostRequest struct {
	DTO
} //@name Binary.PostRequest

type PostResponse struct {
	dto.BaseDTO
	DTO
} //@name Binary.PostResponse

// ***************************
// * PUT
// ***************************
type PutRequest struct {
	ID int `json:"id" example:"1"`
	DTO
} //@name Binary.PutRequest

type PutResponse struct {
	dto.BaseDTO
	DTO
} //@name Binary.PutResponse

// ***************************
// * GET
// ***************************
type GetResponse struct {
	dto.BaseDTO
	DTO
} //@name Binary.GetResponse

// ***************************
// * DELETE
// ***************************
type DeleteRequest struct {
	ID int `json:"id" example:"1"`
} //@name Binary.DeleteRequest

type DeleleResponse struct {
} //@name Binary.DeleteResponse
