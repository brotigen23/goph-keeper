package repository

import (
	"context"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type Users interface {
	Create(ctx context.Context, login, password string) (*model.User, error)

	GetByID(context.Context, int) (*model.User, error)
	GetByLogin(context.Context, string) (*model.User, error)

	Update(context.Context, model.User) (*model.User, error)

	DeleteByID(context.Context, int) (*model.User, error)
}

type Accounts interface {
	Create(ctx context.Context, userID int, login, password string) (*model.AccountData, error)

	GetByID(ctx context.Context, id int) (*model.AccountData, error)
	GetByUserID(ctx context.Context, userID int) ([]model.AccountData, error)

	Update(context.Context, model.AccountData) (*model.AccountData, error)

	DeleteByID(context.Context, int) (*model.AccountData, error)
}

type Text interface {
	Create(ctx context.Context, userID int, data string) (*model.User, error)

	GetByID(ctx context.Context, id int) (*model.TextData, error)
	GetByUserID(ctx context.Context, userID int) ([]model.TextData, error)

	Update(context.Context, model.TextData) (*model.TextData, error)

	DeleteByID(context.Context, int) (*model.TextData, error)
}

type Binary interface {
	Create(ctx context.Context, userID int, data []byte) (*model.BinaryData, error)

	GetByID(ctx context.Context, id int) (*model.BinaryData, error)
	GetByUserID(ctx context.Context, userID int) ([]model.BinaryData, error)

	Update(context.Context, model.BinaryData) (*model.BinaryData, error)

	DeleteByID(context.Context, int) (*model.BinaryData, error)
}

type Cards interface {
	Create(ctx context.Context,
		userID int, number, cardholderName string,
		expireAt time.Time, cvv string) (*model.CardData, error)

	GetByID(ctx context.Context, id int) (*model.TextData, error)
	GetByUserID(ctx context.Context, userID int) ([]model.TextData, error)

	Update(context.Context, model.TextData) (*model.TextData, error)

	DeleteByID(context.Context, int) (*model.TextData, error)
}

type Metadata interface {
	Create(ctx context.Context, tableName string, rowID int, data string) (*model.Metadata, error)

	GetByID(ctx context.Context, id int) (*model.Metadata, error)
	GetByRowID(ctx context.Context, tableName string, rowID int) (*model.Metadata, error)

	Update(context.Context, model.Metadata) (*model.Metadata, error)

	DeleteByID(context.Context, int) (*model.Metadata, error)
}
