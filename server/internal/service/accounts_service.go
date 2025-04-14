package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type AccountsService struct {
	repo repository.Data[model.AccountData]
}

func NewAccountsService(repo repository.Data[model.AccountData]) *AccountsService {
	return &AccountsService{
		repo: repo,
	}
}

func (s AccountsService) Create(ctx context.Context, userID int, login, password string) (*model.AccountData, error) {
	toSave := model.AccountData{UserID: userID, Login: login, Password: password}
	data, err := s.repo.Create(ctx, toSave)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s AccountsService) Update(ctx context.Context, id, userID int, login, password string) (*model.AccountData, error) {
	old, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if old.UserID != userID {
		return nil, ErrIncorrectUserID
	}
	toUpdate := model.AccountData{
		ID:        id,
		Login:     login,
		Password:  password,
		CreatedAt: old.CreatedAt,
	}
	new, err := s.repo.Update(ctx, toUpdate)
	if err != nil {
		return nil, err
	}

	return new, nil

}
func (s AccountsService) GetByUserID(ctx context.Context, userID int) ([]model.AccountData, error) {
	data, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
