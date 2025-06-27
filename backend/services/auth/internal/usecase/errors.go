package usecase

import "errors"

// Create
var (
	ErrBadLogin          = errors.New("bad login")
	ErrBadPassword       = errors.New("bad password")
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)
