//go:build wireinject
// +build wireinject

package http

import (
	"github.com/brotigen23/goph-keeper/auth/internal/infrastructure/repository/memory"
	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/jwt"
	"github.com/brotigen23/goph-keeper/auth/internal/usecase"
	"github.com/google/wire"
)

var DefaultJWTKey = wire.Value("secret")

var DefaultJWTOptions = wire.Value([]jwt.Option{})

func ProviteJWTService(key string, options ...jwt.Option) jwt.Service {
	return jwt.New(key, options...)
}
func WireInitHandler() *handler {
	wire.Build(

		// repo
		memory.New,
		// usecases
		usecase.NewCreateUserUseCase,
		usecase.NewVerifyUserUseCase,
		// controllers
		NewController,
		DefaultJWTKey,
		DefaultJWTOptions,
		ProviteJWTService,

		newHandler,
	)
	return nil
}
