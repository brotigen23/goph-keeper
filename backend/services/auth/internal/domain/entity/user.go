package entity

import "context"

type User struct {
	ID int

	Login    string
	Password string
}

type Updates struct {
	Login    *string
	Password *string
}

// type Filter struct{ Login *string; Password *string}

type Repository interface {
	Create(context.Context, *User) error
	Update(context.Context, Updates) (*User, error)
	GetByID(context.Context, int) (*User, error)
	//Delete(context.Context)
}
