package accountdto

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto"
)

// ***************************
// * Model
// ***************************
type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`

	Metadata string `json:"metadata"`
}

// ***************************
// * POST
// ***************************
type PostRequest struct {
	Account
} //@name Account.PostRequest

func (r *PostRequest) Map(account domain.Account) {
	r.Account = Account{
		Login:    account.Login,
		Password: account.Password,
		Metadata: account.Metadata,
	}
}

type PostResponse struct {
	dto.BaseDTO
	Account
} //@name Account.PostResponse

func (r *PostResponse) Map(account domain.Account) {
	r.BaseDTO = dto.BaseDTO{
		ID:        account.ID,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
	r.Account = Account{
		Login:    account.Login,
		Password: account.Password,
		Metadata: account.Metadata,
	}
}

// ***************************
// * PUT
// ***************************
type PutRequest struct {
	ID int `json:"id"`
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
	ID int `json:"id"`
} //@name Account.DeleteRequest

type DeleleResponse struct {
} //@name Account.DeleteResponse
