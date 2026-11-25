package domain

import "errors"

// Repository
var (
	ErrUserExists   = errors.New("user exists")
	ErrUserNotFound = errors.New("user not found")
)
