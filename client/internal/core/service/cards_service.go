package service

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/cardsdto"
)

type Cards struct {
	client *api.RESTClient
}

func NewCards(client *api.RESTClient) *Cards {
	return &Cards{
		client: client,
	}
}

func (s *Cards) Create(card cardsdto.PostRequest) (*domain.CardData, error) {
	ret := &domain.CardData{}
	cardDTO, err := s.client.PostCard(card)
	if err != nil {
		return nil, err
	}
	ret.Number = cardDTO.Number
	ret.CardholderName = cardDTO.CardholderName
	ret.Expire = cardDTO.Exipre
	ret.CVV = cardDTO.CVV
	ret.Metadata = cardDTO.Metadata

	ret.ID = cardDTO.ID
	ret.CreatedAt = cardDTO.CreatedAt
	ret.UpdatedAt = cardDTO.UpdatedAt
	return ret, nil
}

func (s Cards) Fetch() ([]domain.CardData, error) {
	ret := []domain.CardData{}
	cardsDTO, err := s.client.GetCards()
	if err != nil {
		return nil, err
	}
	// TODO: Вынести в отедльный пакет если можно
	for _, v := range cardsDTO {
		ret = append(ret, domain.CardData{
			Base: domain.Base{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			BaseData: domain.BaseData{
				Metadata: v.Metadata,
			},

			Number:         v.Number,
			CardholderName: v.CardholderName,
			Expire:         v.Exipre,
			CVV:            v.CVV})
	}
	return ret, nil
}

func (s *Cards) Update(card cardsdto.PutRequest) error {
	err := s.client.PutCard(card)
	if err != nil {
		return err
	}
	return nil
}

func (s *Cards) Delete(card cardsdto.DeleteRequest) error {
	_, err := s.client.DeleteCard(card)
	if err != nil {
		return err
	}
	return nil
}
