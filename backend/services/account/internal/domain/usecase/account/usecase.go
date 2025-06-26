package account

import (
	"context"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/repo"
	"github.com/brotigen23/goph-keeper/accounts/internal/domain/usecase"
)

// Usecase service
type Usecase struct {
	repo repo.Repository
}

// Constructor
func New(repo repo.Repository) usecase.Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Create new account data
// TODO: сделать шифрование данных
func (u *Usecase) Create(ctx context.Context, input usecase.Account) (*usecase.AccountWithID, error) {
	account := input.ToEntity()

	err := u.repo.Create(ctx, account)
	if err != nil {
		return nil, translateDBErr(err)
	}
	ret := &usecase.AccountWithID{}
	ret.FromEntity(*account)
	return ret, nil
}

func (u *Usecase) List(ctx context.Context, filter usecase.ListFilter) ([]usecase.AccountWithID, error) {
	f := repo.Filter{
		Login: filter.Login,
	}
	accounts, err := u.repo.Get(ctx, f)
	if err != nil {
		return nil, translateDBErr(err)
	}
	ret := make([]usecase.AccountWithID, 0, len(accounts))
	for _, v := range accounts {
		acc := &usecase.AccountWithID{}
		acc.FromEntity(v)
		ret = append(ret, *acc)
	}

	return ret, nil
}

func (u *Usecase) Update(context.Context, int, usecase.UpdateInput) (*usecase.UpdateOutput, error) {

	return nil, nil
}

func (u *Usecase) Delete(context.Context, int) (*usecase.DeleteOutput, error) {
	return nil, nil
}
