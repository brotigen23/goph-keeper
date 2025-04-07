package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/crypt"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) Create(ctx context.Context, login, password string) (*model.User, error) {
	// Hash password
	passHash, err := crypt.HashPassword(password)
	if err != nil {
		return nil, err
	}
	// Create
	user, err := s.repo.Create(context.Background(), login, passHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	user, err := s.repo.GetByLogin(context.Background(), login)

	switch err {
	case nil:
		return user, nil
	case repository.ErrUserNotFound:
		return nil, ErrUserNotFound
	default:
		return nil, err
	}
}
