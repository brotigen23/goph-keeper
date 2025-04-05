package postgres

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type AccountsRepository struct{}

func (r AccountsRepository) Create(ctx context.Context, userID int, login, password string) (*model.AccountData, error) {
	return nil, nil
}

func (r AccountsRepository) GetByID(ctx context.Context, id int) (*model.AccountData, error) {
	return nil, nil
}
func (r AccountsRepository) GetByUserID(ctx context.Context, userID int) ([]model.AccountData, error) {
	return nil, nil
}

func (r AccountsRepository) Update(context.Context, model.AccountData) (*model.AccountData, error) {
	return nil, nil
}

func (r AccountsRepository) DeleteByID(context.Context, int) (*model.AccountData, error) {
	return nil, nil
}
