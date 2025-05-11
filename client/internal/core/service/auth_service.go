package service

import (
	"os"

	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
)

type AuthService struct {
	client *api.RESTClient

	logger *logger.Logger
}

func NewAuth(client *api.RESTClient) *AuthService {
	return &AuthService{
		client: client,

		logger: logger.New().Testing(),
	}
}

func (s AuthService) GetJWT() string {
	// TODO: crypt or some
	return s.client.GetJWT()
}

func (s AuthService) Register(login, password string) error {
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

func (s AuthService) Login(login, password string) error {
	// TODO: set env jwt
	s.logger.Info("sign in", "login", login, "password", password)
	err := s.client.Login(login, password)
	// If some error
	if err != nil {
		s.logger.Error(err)
		return err
	}
	err = saveTokenToFile(s.client.GetJWT(), ".token")
	return nil
}

func saveTokenToFile(token, path string) error {
	file, err := os.OpenFile(path, 0, 0600)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(token))
	return err

}
