package usecase

import (
	"context"

	"github.com/brotigen23/goph-keeper/auth/internal/domain"
	"github.com/brotigen23/goph-keeper/shared/pkg/crypt"
)

type CreateUserInput struct {
	Login    string
	Password string
}

func (i *CreateUserInput) Validate() error {
	if i.Login == "" {
		return ErrEmptyLogin
	}
	if i.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}

type CreateUserOutput struct {
	ID int
}

type CreateUserUseCase struct {
	repo domain.Repository
}

func NewCreateUserUseCase(repo domain.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	pass, err := crypt.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Login:    input.Login,
		Password: pass,
	}
	err = u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	ret := &CreateUserOutput{
		ID: user.ID,
	}
	return ret, nil
}
