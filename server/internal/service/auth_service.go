package service

import (
	"context"

	"github.com/brotigen23/goph-keeper/server/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) error
}
