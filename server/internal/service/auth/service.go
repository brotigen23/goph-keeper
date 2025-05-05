package auth

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/crypt"
)

type Service struct {
	repo repository.User
}

func New(repo repository.User) service.AuthService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, user *model.User) error {
	// Hash password
	hashPassword, err := crypt.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	// Create
	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(ctx context.Context, user *model.User) error {
	userFromDB, err := s.repo.GetByLogin(context.Background(), user.Login)
	if err != nil {
		return service.TranslateErr(err)
	}
	if err = crypt.CheckPasswordHash(user.Password, userFromDB.Password); err != nil {
		return service.ErrPasswordIncorrect
	}
	*user = *userFromDB
	return nil
}
