package postgres

import (
	"database/sql"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u Users) Create(model.User) (*model.User, error) {
	return nil, nil
}

func (u Users) GetByID(id int) (*model.User, error) {
	return nil, nil
}
func (u Users) GetByLogin(login string) (*model.User, error) {
	return nil, nil
}

func (u Users) Update(model.User) (*model.User, error) {
	return nil, nil
}

func (u Users) Delete(model.User) (*model.User, error) {
	return nil, nil
}
