package usecase

import (
	"context"

	"github.com/brotigen23/goph-keeper/auth/internal/domain"
	"github.com/brotigen23/goph-keeper/shared/pkg/crypt"
)

type VerifyUserInput struct {
	Login    string
	Password string
}

func (i *VerifyUserInput) Validate() error {
	if i.Login == "" {
		return ErrBadLogin
	}
	if i.Password == "" {
		return ErrBadPassword
	}
	return nil
}

type VerifyUserOutput struct {
	ID int
}

type VerifyUserUseCase struct {
	repo domain.Repository
}

func NewVerifyUserUseCase(repo domain.Repository) *VerifyUserUseCase {
	return &VerifyUserUseCase{
		repo: repo,
	}
}

func (u *VerifyUserUseCase) Execute(ctx context.Context, input VerifyUserInput) (*VerifyUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	filter := domain.Filter{
		Login: &input.Login,
	}
	user, err := u.repo.Get(ctx, filter)
	if err != nil {
		return nil, ErrUserNotFound
	}
	err = crypt.CheckPasswordHash(input.Password, user.Password)
	if err != nil {
		return nil, ErrPasswordIncorrect
	}

	ret := &VerifyUserOutput{
		ID: user.ID,
	}

	return ret, nil
}
