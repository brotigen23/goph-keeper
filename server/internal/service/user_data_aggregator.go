package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/crypt"
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
	user, err := a.userService.repo.GetByLogin(ctx, login)
	switch err {
	case nil:
		break
	case repository.ErrUserNotFound:
		return nil, ErrUserNotFound
	default:
		return nil, err
	}
	err = crypt.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, ErrIncorrectPassword
	}
	return user, nil
}

func (a UserDataAggregator) CreateAccount(ctx context.Context,
	userID int, login, password string) (*model.AccountData, error) {
	ret, err := a.accountsService.Create(ctx, userID, login, password)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (a UserDataAggregator) GetAccounts(ctx context.Context,
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

func (a UserDataAggregator) UpdateAccount(ctx context.Context,
	userID, id int, login, password *string) (*model.AccountData, error) {
	// Update Account
	oldAccount, err := a.accountsService.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userID != oldAccount.UserID {
		return nil, ErrIncorrectUserID
	}
	if login != nil {
		oldAccount.Login = *login
	}

	newAccount, err := a.accountsService.repo.Update(ctx, *oldAccount)
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

func (a UserDataAggregator) UpdateMetadata(ctx context.Context,
	id int, metadata string) (*model.Metadata, error) {
	newMetadata := model.Metadata{ID: id, Data: metadata}
	savedMetadata, err := a.metadataService.repo.Update(ctx, newMetadata)
	if err != nil {
		return nil, err
	}
	return savedMetadata, nil
}
