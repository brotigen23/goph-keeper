package handler

import "errors"

var (
	ErrIncorrectPassword = errors.New("Incorrect password")
	ErrUserNotFound      = errors.New("User not found")
)
