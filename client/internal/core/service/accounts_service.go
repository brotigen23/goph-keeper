package service

import (
	"os"

	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
)

type AccountsService struct {
	client api.APIClient
}

func NewAccounts(client api.APIClient) *AccountsService {
	return &AccountsService{
		client: client,
	}
}

func (s AccountsService) GetAccounts() ([]domain.AccountData, error) {
	jwt := os.Getenv("KEEPER_JWT")
	s.client.SetJWT(jwt)
	accounts, err := s.client.GetAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s AccountsService) CreateAccount(account domain.AccountData) error {
	return nil
}
