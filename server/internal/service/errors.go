package service

import "errors"

var (
	ErrUserExists   = errors.New("User already exist")
	ErrUserNotFound = errors.New("User not found")

	ErrIncorrectPassword = errors.New("Incorrect password")
	ErrIncorrectUserID   = errors.New("Incorrect user id")

	ErrDataNotFound = errors.New("Data not found")
)
