package textdto

import "github.com/brotigen23/goph-keeper/client/internal/core/dto"

// ***************************
// * Model
// ***************************
type DTO struct {
	Data string `json:"data" example:"user"`

	Metadata string `json:"metadata" example:"metadata"`
}

// ***************************
// * POST
// ***************************
type PostRequest struct {
	DTO
} //@name Text.PostRequest

type PostResponse struct {
	dto.BaseDTO
	DTO
} //@name Text.PostResponse

// ***************************
// * PUT
// ***************************
type PutRequest struct {
	ID int `json:"id" example:"1"`
	DTO
} //@name Text.PutRequest

type PutResponse struct {
	dto.BaseDTO
	DTO
} //@name Text.PutResponse

// ***************************
// * GET
// ***************************
type GetResponse struct {
	dto.BaseDTO
	DTO
} //@name Text.GetResponse

// ***************************
// * DELETE
// ***************************
type DeleteRequest struct {
	ID int `json:"id" example:"1"`
} //@name Text.DeleteRequest

type DeleleResponse struct {
} //@name Text.DeleteResponse
