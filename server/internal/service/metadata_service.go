package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type MetadataService struct {
	repo repository.Metadata
}

func NewMetadataService(repo repository.Metadata) *MetadataService {
	return &MetadataService{
		repo: repo,
	}
}

func (s MetadataService) GetByID(ctx context.Context, id int) (*model.Metadata, error) {
	ret, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s MetadataService) Update(ctx context.Context, id int, data string) (*model.Metadata, error) {
	toSave := &model.Metadata{ID: id, Data: data}
	ret, err := s.repo.Update(ctx, *toSave)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s MetadataService) GetAccount(ctx context.Context, rowID int) (*model.Metadata, error) {
	return nil, nil
}
