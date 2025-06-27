//go:build wireinject
// +build wireinject

package http

import (
	"github.com/brotigen23/goph-keeper/auth/internal/infrastructure/repository/memory"
	"github.com/brotigen23/goph-keeper/auth/internal/usecase"
	"github.com/google/wire"
)

func WireInitHandler() *handler {
	wire.Build(
		// repo
		memory.New,
		// usecases
		usecase.NewCreateUserUseCase,
		usecase.NewVerifyUserUseCase,
		// controllers
		NewController,

		newHandler,
	)
	return nil
}
