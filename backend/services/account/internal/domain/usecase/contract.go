package usecase

import (
	"context"
)

type Usecase interface {
	Create(context.Context, Account) (*AccountWithID, error)
	Update(context.Context, int, UpdateInput) (*UpdateOutput, error)
	List(context.Context, ListFilter) ([]AccountWithID, error)
	Delete(context.Context, int) (*DeleteOutput, error)
}
