package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type UserDataAggregator struct {
	userService       *UserService
	accountsService   *AccountsService
	textDataService   *TextDataService
	binaryDataService *BinaryDataService
	cardsService      *CardsService
	metadataService   *MetadataService
}

func NewAggregator(
	u *UserService,
	a *AccountsService,
	t *TextDataService,
	b *BinaryDataService,
	c *CardsService,
	m *MetadataService,
) *UserDataAggregator {
	return &UserDataAggregator{
		userService:       u,
		accountsService:   a,
		textDataService:   t,
		binaryDataService: b,
		cardsService:      c,
		metadataService:   m,
	}
}

func (a UserDataAggregator) CreateNewUser(ctx context.Context, login, password string) (*model.User, error) {
	user, err := a.userService.GetUserByLogin(ctx, login)
	switch err {
	case nil:
		return user, ErrUserExists
	case ErrUserNotFound:
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

func (a UserDataAggregator) ValidateUserLogin(ctx context.Context, login, password string) (*model.User, error) {
	return nil, nil
}

func (a UserDataAggregator) GetUserAccountsData() {

}
