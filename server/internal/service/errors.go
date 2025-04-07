package service

import "errors"

var (
	ErrUserExists   = errors.New("User already exist")
	ErrUserNotFound = errors.New("User not found")
)
