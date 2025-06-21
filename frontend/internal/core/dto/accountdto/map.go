package accountdto

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto"
)

func baseDataToDTO(account domain.Account) dto.BaseDTO {
	return dto.BaseDTO{
		ID:        account.ID,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func accountToDTO(account domain.Account) Account {
	return Account{
		Login:    account.Login,
		Password: account.Password,
		Metadata: account.Metadata,
	}
}

func (r *PostRequest) Map(account domain.Account) {
	r.Account = accountToDTO(account)
}

func (r *PostResponse) Map(account domain.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}

func (r *PutRequest) Map(account domain.Account) {
	r.ID = account.ID
	r.Account = accountToDTO(account)
}

func (r *PutResponse) Map(account domain.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}

func (r *GetResponse) Map(account domain.Account) {
	r.BaseDTO = baseDataToDTO(account)
	r.Account = accountToDTO(account)
}
