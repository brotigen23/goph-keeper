package service

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/binarydto"
)

type Binary struct {
	client *api.RESTClient
}

func NewBinary(client *api.RESTClient) *Binary {
	return &Binary{
		client: client,
	}
}

func (s *Binary) Create(binary binarydto.PostRequest) (*domain.BinaryData, error) {
	ret := &domain.BinaryData{}
	binaryDTO, err := s.client.PostBinary(binary)
	if err != nil {
		return nil, err
	}
	ret.Data = binaryDTO.Data
	ret.Metadata = binaryDTO.Metadata

	ret.ID = binaryDTO.ID
	ret.CreatedAt = binaryDTO.CreatedAt
	ret.UpdatedAt = binaryDTO.UpdatedAt
	return ret, nil
}

func (s Binary) Fetch() ([]domain.BinaryData, error) {
	ret := []domain.BinaryData{}
	binaryDTO, err := s.client.GetBinary()
	if err != nil {
		return nil, err
	}
	// TODO: Вынести в отедльный пакет если можно
	for _, v := range binaryDTO {
		ret = append(ret, domain.BinaryData{
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

func (s *Binary) Update(binary binarydto.PutRequest) error {
	err := s.client.PutBinary(binary)
	if err != nil {
		return err
	}
	return nil
}

func (s *Binary) Delete(binary binarydto.DeleteRequest) error {
	_, err := s.client.DeleteBinary(binary)
	if err != nil {
		return err
	}
	return nil
}
