package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) Create(login, password string) (*model.User, error) {
	// Check if user already exists
	existUser, err := s.GetUserByLogin(login)
	switch err {
	case nil:
		return existUser, ErrUserExists
	case repository.ErrUserNotFound:
		break
	default:
		return nil, err
	}
	// Hash password

	// Create
	user, err := s.repo.Create(context.Background(), login, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) GetUserByLogin(login string) (*model.User, error) {

	return nil, nil
}
