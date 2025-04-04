package repository

import "github.com/brotigen23/goph-keeper/server/internal/model"

type Users interface {
	Create(model.User) (*model.User, error)

	GetByID(id int) (*model.User, error)
	GetByLogin(login string) (*model.User, error)

	Update(model.User) (*model.User, error)

	Delete(model.User) (*model.User, error)
}
