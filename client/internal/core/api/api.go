package api

import "github.com/brotigen23/goph-keeper/client/internal/core/domain"

type APIClient interface {
	SetJWT(string)
	GetJWT() string

	Register(string, string) error
	Login(string, string) error
	GetAccounts() ([]domain.AccountData, error)
}
