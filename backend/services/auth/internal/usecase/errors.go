package usecase

import "errors"

// Create
var (
	ErrEmptyLogin    = errors.New("login is empty")
	ErrEmptyPassword = errors.New("password is empty")

	ErrUserExist = errors.New("user exist")
)

// Login
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)
