package accountdto

import (
	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/model"
)

func baseDataToDTO(account model.Account) dto.BaseDTO {
	return dto.BaseDTO{
		ID:        account.ID,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func accountToDTO(account model.Account) Account {
	return Account{
		Login:    account.Login,
		Password: account.Password,
		Metadata: account.Metadata,
	}
}

func (r *PostRequest) Map(account model.Account) {
	r.Account = accountToDTO(account)
}

func (r *PostResponse) Map(account model.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}

func (r *PutRequest) Map(account model.Account) {
	r.ID = account.ID
	r.Account = accountToDTO(account)
}

func (r *PutResponse) Map(account model.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}

func (r *GetResponse) Map(account model.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}
