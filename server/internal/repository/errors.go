package repository

import "errors"

var (
	// Users
	ErrUserExists   = errors.New("User already exists")
	ErrUserNotFound = errors.New("User not found")

	ErrAccountsDataNotFound = errors.New("Accounts data not found")
	ErrTextDataNotFound     = errors.New("Text data not found")
	ErrBinaryDataNotFound   = errors.New("Binary data not found")
	ErrCardsDataNotFound    = errors.New("Card data not found")

	ErrMetadataNotFound = errors.New("Metadata not found")
)
