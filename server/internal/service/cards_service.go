package service

import (
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type CardsService struct {
}

func NewCardsService(repo repository.Data[model.CardData]) *CardsService {
	return nil
}
