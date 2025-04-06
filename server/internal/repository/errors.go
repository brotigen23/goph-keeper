package repository

import "errors"

var (
	// Users
	ErrUserExists   = errors.New("User already exists")
	ErrUserNotFound = errors.New("User not found")
)
