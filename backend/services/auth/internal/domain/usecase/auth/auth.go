package auth

import (
	"context"

	"github.com/brotigen23/goph-keeper/auth/internal/domain/entity"
	"github.com/brotigen23/goph-keeper/auth/internal/domain/usecase"
	"github.com/brotigen23/goph-keeper/shared/pkg/crypt"
)

type Auth struct {
	repo entity.Repository
}

func New(repo entity.Repository) usecase.Auth {
	return &Auth{
		repo: repo,
	}
}

func (a *Auth) CreateUser(ctx context.Context, user usecase.User) (*usecase.CreatedUser, error) {
	// pass, err := crypt.Hash(u.Password)
	// user.Password = pass
	u := user.ToEntity()
	err := a.repo.Create(ctx, u)
	if err != nil {
		return nil, translateRepoErr(err)
	}
	var ret *usecase.CreatedUser
	ret.FromEntity(*u)
	return ret, nil
}

func (a *Auth) VerifyUser(ctx context.Context, input usecase.User) error {
	user := input.ToEntity()
	u, err := a.repo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	err = crypt.CheckHash(user.Password, u.Password)
	if err != nil {
		return err
	}

	return nil
}
