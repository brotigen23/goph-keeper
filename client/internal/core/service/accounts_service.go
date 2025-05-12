package service

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
)

type Accounts struct {
	client *api.RESTClient
}

func NewAccounts(client *api.RESTClient) *Accounts {
	return &Accounts{
		client: client,
	}
}

func (s *Accounts) Create(account accountdto.PostRequest) (*domain.Account, error) {
	ret := &domain.Account{}
	accountDTO, err := s.client.PostAccount(account)
	if err != nil {
		return nil, err
	}
	ret.Login = accountDTO.Login
	ret.Password = accountDTO.Password
	ret.Metadata = accountDTO.Metadata

	ret.ID = accountDTO.ID
	ret.CreatedAt = accountDTO.CreatedAt
	ret.UpdatedAt = accountDTO.UpdatedAt
	return ret, nil
}

func (s Accounts) Fetch() ([]domain.Account, error) {
	ret := []domain.Account{}
	accountsDTO, err := s.client.GetAccounts()
	if err != nil {
		return nil, err
	}
	// TODO: Вынести в отедльный пакет если можно
	for _, v := range accountsDTO {
		ret = append(ret, domain.Account{
			Base: domain.Base{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			BaseData: domain.BaseData{
				Metadata: v.Metadata,
			},
			Login:    v.Login,
			Password: v.Password,
		})
	}
	return ret, nil
}

func (s *Accounts) Update(account accountdto.PutRequest) error {
	err := s.client.PutAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func (s *Accounts) Delete(account accountdto.DeleteRequest) error {
	_, err := s.client.DeleteAccount(account)
	if err != nil {
		return err
	}
	return nil
}
