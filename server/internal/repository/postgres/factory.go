package postgres

import (
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres/account"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres/user"
	"github.com/jmoiron/sqlx"
)

type Factory struct {
	db *sqlx.DB
}

func NewFactory(db *sqlx.DB) repository.Factory {
	return &Factory{
		db: db,
	}
}

func (f *Factory) NewUserRepository() repository.User {
	return user.New(f.db)
}

func (f *Factory) NewAccountRepository() repository.AccountsData {
	return account.New(f.db)
}
