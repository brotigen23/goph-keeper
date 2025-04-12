package repository

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type Users interface {
	Create(ctx context.Context, login, password string) (*model.User, error)

	GetByID(context.Context, int) (*model.User, error)
	GetByLogin(context.Context, string) (*model.User, error)

	Update(context.Context, model.User) (*model.User, error)

	DeleteByID(context.Context, int) (*model.User, error)
}

type Data[T model.Model] interface {
	Create(context.Context, T) (*T, error)

	GetByID(context.Context, int) (*T, error)
	GetByUserID(context.Context, int) ([]T, error)

	Update(context.Context, T) (*T, error)

	DeleteByID(context.Context, int) (*T, error)
}

type Metadata interface {
	Create(ctx context.Context, data string) (*model.Metadata, error)

	GetByID(ctx context.Context, id int) (*model.Metadata, error)

	Update(context.Context, model.Metadata) (*model.Metadata, error)

	DeleteByID(context.Context, int) (*model.Metadata, error)
}
