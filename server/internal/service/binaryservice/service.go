package binaryservice

import (
	"context"
	"errors"
	"log"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/internal/service"
)

// TODO: log
type Service struct {
	repo repository.Binary
}

func New(repo repository.Binary) service.BinaryService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, item *model.BinaryData) error {
	err := s.repo.Create(ctx, item)
	return err
}

func (s *Service) Update(ctx context.Context, item *model.BinaryData) error {
	err := s.repo.Update(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id int) (*model.BinaryData, error) {
	// Пока не используется
	return nil, service.ErrNotImplement
}

func (s *Service) Delete(ctx context.Context, userID, id int) error {
	item, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	log.Println(item)
	log.Println(userID)
	if item.UserID != userID {
		return errors.New("own err")
	}
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserData(ctx context.Context, userID int) ([]model.BinaryData, error) {
	ret, err := s.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
