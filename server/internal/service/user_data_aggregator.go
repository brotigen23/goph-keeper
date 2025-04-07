package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
)

type UserDataAggregator struct {
	// Servicies
	userService *UserService
}

func NewAggregator() *UserDataAggregator {
	return &UserDataAggregator{}
}

func (a UserDataAggregator) CreateNewUser(ctx context.Context, login, password string) (*model.User, error) {
	user, err := a.userService.GetUserByLogin(ctx, login)
	switch err {
	case nil:
		return user, ErrUserExists
	case repository.ErrUserNotFound:
		break
	default:
		return nil, err
	}
	user, err = a.userService.Create(ctx, login, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a UserDataAggregator) GetUserAccountsData() {

}
