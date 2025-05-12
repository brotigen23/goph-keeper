package service

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
)

type Auth struct {
	client *api.RESTClient

	logger *logger.Logger
}

func NewAuth(client *api.RESTClient) *Auth {
	return &Auth{
		client: client,

		logger: logger.New().Testing(),
	}
}

func (s Auth) GetJWT() string {
	// TODO: crypt or some
	return s.client.GetJWT()
}

func (s Auth) Register(login, password string) error {
	// TODO: set env jwt
	s.logger.Info("sign in", "login", login, "password", password)
	err := s.client.Register(login, password)
	// If some error
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}

func (s Auth) Login(login, password string) error {
	// TODO: set env jwt
	s.logger.Info("sign in", "login", login, "password", password)
	err := s.client.Login(login, password)
	// If some error
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
