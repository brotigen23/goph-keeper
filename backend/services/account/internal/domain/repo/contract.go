package repo

import (
	"context"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/entity"
)

type Filter struct {
	Login    *string
	Password *string
}

type Updates struct {
	ID int

	Login    *string
	Password *string
}

type Repository interface {
	Create(context.Context, *entity.Account) error
	Update(context.Context, Updates) (*entity.Account, error)
	Get(context.Context, Filter) ([]entity.Account, error)
	Delete(context.Context, int) (*entity.Account, error)
}
