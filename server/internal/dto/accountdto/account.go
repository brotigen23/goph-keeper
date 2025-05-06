package accountdto

import "github.com/brotigen23/goph-keeper/server/internal/dto"

// ***************************
// * Model
// ***************************
type Account struct {
	Login    string `json:"login" example:"user"`
	Password string `json:"password" example:"user"`

	Metadata string `json:"metadata" example:"metadata"`
}

// ***************************
// * POST
// ***************************
type PostRequest struct {
	Account
} //@name Account.PostRequest

type PostResponse struct {
	dto.BaseDTO
	Account
} //@name Account.PostResponse

// ***************************
// * PUT
// ***************************
type PutRequest struct {
	ID int `json:"id" example:"1"`
	Account
} //@name Account.PutRequest

type PutResponse struct {
	dto.BaseDTO
	Account
} //@name Account.PutResponse

// ***************************
// * GET
// ***************************
type GetResponse struct {
	dto.BaseDTO
	Account
} //@name Account.Get.Response

// ***************************
// * DELETE
// ***************************
type DeleteRequest struct {
	ID int `json:"id" example:"1"`
} //@name Account.DeleteRequest

type DeleleResponse struct {
} //@name Account.DeleteResponse
