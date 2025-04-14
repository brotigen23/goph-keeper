package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type BinaryDataService struct {
	repo repository.Data[model.BinaryData]
}

func NewBinaryDataService(repo repository.Data[model.BinaryData]) *BinaryDataService {
	return &BinaryDataService{
		repo: repo,
	}
}

func (s BinaryDataService) Create(ctx context.Context, userID int, data []byte) (*model.BinaryData, error) {
	toSave := model.BinaryData{UserID: userID, Data: data}
	savedData, err := s.repo.Create(ctx, toSave)
	if err != nil {
		return nil, err
	}
	return savedData, nil
}
