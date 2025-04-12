package service

import (
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type TextDataService struct {
}

func NewTextDataService(repo repository.Data[model.TextData]) *TextDataService {
	return nil
}
