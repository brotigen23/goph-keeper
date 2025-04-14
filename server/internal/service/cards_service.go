package service

import (
	"context"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type CardsService struct {
	repo repository.Data[model.CardData]
}

func NewCardsService(repo repository.Data[model.CardData]) *CardsService {
	return &CardsService{
		repo: repo,
	}
}

func (s CardsService) Create(ctx context.Context,
	userID int,
	number, cardholderName string,
	expire time.Time,
	cvv string) (*model.CardData, error) {
	toSave := model.CardData{UserID: userID, Number: number, CardholderName: cardholderName, Expire: expire, CVV: cvv}
	savedData, err := s.repo.Create(ctx, toSave)
	if err != nil {
		return nil, err
	}
	return savedData, nil
}
