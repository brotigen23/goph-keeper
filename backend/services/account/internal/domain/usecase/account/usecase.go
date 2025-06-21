package account

import (
	"context"

	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/entity"
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/repo"
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/usecase"
)

// Usecase service
type Usecase struct {
	repo repo.Repository
}

// Constructor
func New(repo repo.Repository) *Usecase {
	return &Usecase{}
}

// Create new account data
func (u *Usecase) Create(ctx context.Context, input usecase.Account) (*usecase.AccountWithID, error) {
	account := input.ToEntity()
	err := u.repo.Create(ctx, account)
	if err != nil {
		return nil, translateDBErr(err)
	}
	ret := &usecase.AccountWithID{
		ID: account.ID,
		Account: usecase.Account{
			Login:    account.Login,
			Password: account.Password,
		},
	}
	return ret, nil
}

func (u *Usecase) List(ctx context.Context, filters *usecase.ListFilter) ([]entity.Account, error) {
	return nil, ErrNotImplemented
}
