package dto

import (
	"time"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/entity"
)

type Account struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Login    string `db:"login"`
	Password string `db:"password"`
}

func NewAccountFromEntity(a entity.Account) *Account {
	return &Account{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Login:     a.Login,
		Password:  a.Password,
	}
}

func (a *Account) ToEntity() *entity.Account {
	return &entity.Account{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Login:     a.Login,
		Password:  a.Password,
	}
}
