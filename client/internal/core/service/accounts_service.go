package service

import (
	"os"

	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
)

type AccountsService struct {
	client *api.RESTClient
}

func NewAccounts(client *api.RESTClient) *AccountsService {
	return &AccountsService{
		client: client,
	}
}

// TODO: изменить входной параметр
func (s AccountsService) CreateAccount(account *domain.Account) error {
	request := accountdto.PostRequest{
		Account: accountdto.Account{
			Login:    account.Login,
			Password: account.Password,
			Metadata: account.Metadata,
		},
	}
	response, err := s.client.PostAccount(request)
	if err != nil {
		return nil
	}
	account.ID = response.ID
	account.CreatedAt = response.CreatedAt
	account.UpdatedAt = response.UpdatedAt
	return nil
}

func (s AccountsService) GetAccounts() ([]domain.Account, error) {
	ret := []domain.Account{}
	jwt := os.Getenv("KEEPER_JWT")
	s.client.SetJWT(jwt)
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
