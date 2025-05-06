package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type DataService[T model.Model] interface {
	Create(ctx context.Context, item *T) error
	Update(ctx context.Context, item *T) error
	Get(ctx context.Context, id int) (*T, error)
	Delete(ctx context.Context, userid, id int) error

	GetUserData(ctx context.Context, userID int) ([]T, error)
}

type AccountService interface {
	DataService[model.Account]
}

type TextService interface {
	DataService[model.TextData]
}

type BinaryService interface {
	DataService[model.BinaryData]
}

type CardService interface {
	DataService[model.CardData]
}
