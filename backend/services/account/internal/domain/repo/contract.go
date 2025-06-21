package repo

import (
	"context"

	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/entity"
)

type Filter struct {
	Login    *string
	Password *string
}

type Repository interface {
	Create(context.Context, *entity.Account) error
	Get(context.Context, Filter) ([]entity.Account, error)
}
