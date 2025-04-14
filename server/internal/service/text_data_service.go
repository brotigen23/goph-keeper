package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type TextDataService struct {
	repo repository.Data[model.TextData]
}

func NewTextDataService(repo repository.Data[model.TextData]) *TextDataService {
	return &TextDataService{
		repo: repo,
	}
}

func (s TextDataService) Create(ctx context.Context, userID int, data string) (*model.TextData, error) {
	toSave := model.TextData{UserID: userID, Data: data}
	savedData, err := s.repo.Create(ctx, toSave)
	if err != nil {
		return nil, err
	}
	return savedData, nil
}
