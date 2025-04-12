package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
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
	userRepo repository.Users,
	accountsRepo repository.Data[model.AccountData],
	textDataRepo repository.Data[model.TextData],
	binaryDataRepo repository.Data[model.BinaryData],
	cardsRepo repository.Data[model.CardData],
	metadataRepo repository.Metadata,
) *UserDataAggregator {
	return &UserDataAggregator{
		userService:       NewUserService(userRepo),
		accountsService:   NewAccountsService(accountsRepo),
		textDataService:   NewTextDataService(textDataRepo),
		binaryDataService: NewBinaryDataService(binaryDataRepo),
		cardsService:      NewCardsService(cardsRepo),
		metadataService:   NewMetadataService(metadataRepo),
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

func (a UserDataAggregator) ValidateUserSingIn(ctx context.Context, login, password string) (*model.User, error) {
	return nil, nil
}

func (a UserDataAggregator) CreateUserAccountsData(ctx context.Context,
	userID int, login, password string) (*model.AccountData, string, error) {
	_, err := a.accountsService.Create(ctx, userID, login, password)
	if err != nil {
		return nil, "", err
	}

	return nil, "", nil
}

func (a UserDataAggregator) GetUserAccountsData(ctx context.Context,
	userID int) ([]model.AccountData, []model.Metadata, error) {
	ret, err := a.accountsService.GetByUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}
	retMetadata := []model.Metadata{}
	for _, v := range ret {
		metadata, err := a.metadataService.GetByID(ctx, v.ID)
		if err != nil {
			return nil, nil, err
		}
		retMetadata = append(retMetadata, *metadata)
	}
	return ret, retMetadata, nil
}

func (a UserDataAggregator) CreateNewTextData(ctx context.Context, data string) (*model.TextData, error) {
	return nil, nil
}
