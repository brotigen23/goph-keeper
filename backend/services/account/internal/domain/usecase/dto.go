package usecase

import "github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/entity"

type Account struct {
	Login    string
	Password string
}

func (a *Account) ToEntity() *entity.Account {
	return &entity.Account{
		Login:    a.Login,
		Password: a.Password,
	}
}

type AccountWithID struct {
	ID int
	Account
}

func (a *AccountWithID) ToEntity() *entity.Account {
	return &entity.Account{
		ID:       a.ID,
		Login:    a.Login,
		Password: a.Password,
	}
}

type CreateInput struct {
	Account
}

type CreateOutput struct {
	ID int
	Account
}

type UpdateInput struct {
	Login    *string
	Password *string
}

type UpdateOutput struct {
	ID int
	Account
}

type ListFilter struct {
	Login *string
}

type ListOutput struct {
	Account
}

type DeleteOutput struct {
	Account
}
