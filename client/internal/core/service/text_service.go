package service

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/textdto"
)

type Text struct {
	client *api.RESTClient
}

func NewText(client *api.RESTClient) *Text {
	return &Text{
		client: client,
	}
}

func (s *Text) Create(text textdto.PostRequest) (*domain.TextData, error) {
	ret := &domain.TextData{}
	accountDTO, err := s.client.PostText(text)
	if err != nil {
		return nil, err
	}
	ret.Data = accountDTO.Data
	ret.Metadata = accountDTO.Metadata

	ret.ID = accountDTO.ID
	ret.CreatedAt = accountDTO.CreatedAt
	ret.UpdatedAt = accountDTO.UpdatedAt
	return ret, nil
}

func (s Text) Fetch() ([]domain.TextData, error) {
	ret := []domain.TextData{}
	textDTO, err := s.client.GetText()
	if err != nil {
		return nil, err
	}
	// TODO: Вынести в отедльный пакет если можно
	for _, v := range textDTO {
		ret = append(ret, domain.TextData{
			Base: domain.Base{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			BaseData: domain.BaseData{
				Metadata: v.Metadata,
			},
			Data: v.Data,
		})
	}
	return ret, nil
}

func (s *Text) Update(text textdto.PutRequest) error {
	err := s.client.PutText(text)
	if err != nil {
		return err
	}
	return nil
}

func (s *Text) Delete(text textdto.DeleteRequest) error {
	_, err := s.client.DeleteText(text)
	if err != nil {
		return err
	}
	return nil
}
