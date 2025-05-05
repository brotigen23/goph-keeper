package repository

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

// Создать отдельные сущности если понадобится
type Base[T model.Model] interface {
	Create(context.Context, *T) error
	Get(context.Context, int) (*T, error)
	Update(context.Context, *T) error
	Delete(context.Context, int) error
}
type User interface {
	Base[model.User]
	GetByLogin(context.Context, string) (*model.User, error)
}

// Для пользовательских данных
type UserData[T model.Model] interface {
	Base[T]
	GetAll(context.Context, int) ([]T, error)
}

type AccountsData interface {
	UserData[model.Account]
}
