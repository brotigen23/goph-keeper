package entity

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrLogin              = errors.New("Login must be not empty")
	ErrPassword           = errors.New("Password must be not empty")
)
