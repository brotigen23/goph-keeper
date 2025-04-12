package service

import (
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type BinaryDataService struct {
}

func NewBinaryDataService(repo repository.Data[model.BinaryData]) *BinaryDataService {
	return nil
}
