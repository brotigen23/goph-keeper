package account

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/internal/service"
)

// TODO: log
type Service struct {
	repo repository.AccountsData
}

func New(repo repository.AccountsData) service.AccountService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, item *model.Account) error {
	err := s.repo.Create(ctx, item)
	return err
}

func (s *Service) Update(ctx context.Context, item *model.Account) error {
	err := s.repo.Update(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id int) (*model.Account, error) {
	// Пока не используется
	return nil, service.ErrNotImplement
}

func (s *Service) Delete(ctx context.Context, id int) error {
	// Метка или немедленное удаление
	return service.ErrNotImplement
}

func (s *Service) GetUserData(ctx context.Context, userID int) ([]model.Account, error) {
	ret, err := s.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
