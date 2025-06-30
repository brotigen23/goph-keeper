package http

import (
	"context"

	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/request"
	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/response"
	"github.com/brotigen23/goph-keeper/auth/internal/usecase"
)

type controller struct {
	ucRegister *usecase.CreateUserUseCase
	ucVerify   *usecase.VerifyUserUseCase
}

func NewController(ucRegister *usecase.CreateUserUseCase, ucVerify *usecase.VerifyUserUseCase) *controller {
	return &controller{
		ucRegister: ucRegister,
		ucVerify:   ucVerify,
	}
}

func (c *controller) register(ctx context.Context, req request.Register) (*response.Register, error) {
	input := usecase.CreateUserInput{
		Login:    req.Login,
		Password: req.Password,
	}
	user, err := c.ucRegister.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	ret := &response.Register{
		ID: user.ID,
	}
	return ret, nil
}

func (c *controller) login(ctx context.Context, req request.Login) (*response.Login, error) {
	input := usecase.VerifyUserInput{
		Login:    req.Login,
		Password: req.Password,
	}
	output, err := c.ucVerify.Execute(ctx, input)
	if err != nil {
		return nil, err
	}
	resp := &response.Login{
		ID: output.ID,
	}
	return resp, nil
}
