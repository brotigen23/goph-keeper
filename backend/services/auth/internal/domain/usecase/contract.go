package usecase

import "context"

type Auth interface {
	CreateUser(context.Context, User) (*CreatedUser, error)
	VerifyUser(context.Context, User) error
}
