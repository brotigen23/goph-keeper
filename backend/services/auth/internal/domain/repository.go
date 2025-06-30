package domain

import "context"

type Filter struct {
	ID    *int
	Login *string
}

type Updates struct {
	Login    *string
	Password *string
}

type Repository interface {
	Create(context.Context, *User) error
	Update(context.Context, Updates) (*User, error)
	Get(context.Context, Filter) (*User, error)
	//Delete(context.Context)
}
