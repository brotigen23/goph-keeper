package cardsdto

import (
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
)

// ***************************
// * Model
// ***************************
type DTO struct {
	Number         string    `json:"number" example:"123456789123456"`
	CardholderName string    `json:"cardholder_name" example:"IVAN IVANOV"`
	Exipre         time.Time `json:"expire" example:"2033-01-01T00:00:00.000000Z"`
	CVV            string    `json:"cvv" example:"123"`

	Metadata string `json:"metadata" example:"metadata"`
}

// ***************************
// * POST
// ***************************
type PostRequest struct {
	DTO
} //@name Card.PostRequest

type PostResponse struct {
	dto.BaseDTO
	DTO
} //@name Card.PostResponse

// ***************************
// * PUT
// ***************************
type PutRequest struct {
	ID int `json:"id" example:"1"`
	DTO
} //@name Card.PutRequest

type PutResponse struct {
	dto.BaseDTO
	DTO
} //@name Card.PutResponse

// ***************************
// * GET
// ***************************
type GetResponse struct {
	dto.BaseDTO
	DTO
} //@name Card.GetResponse

// ***************************
// * DELETE
// ***************************
type DeleteRequest struct {
	ID int `json:"id" example:"1"`
} //@name Card.DeleteRequest

type DeleleResponse struct {
} //@name Card.DeleteResponse
