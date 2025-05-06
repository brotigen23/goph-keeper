package repository

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

var (
	ErrNotImplement = errors.New("Not implement")
	ErrNotFound     = errors.New("Not found")
	ErrConflict     = errors.New("Conflict")

	ErrNoUpdate = errors.New("No rows were update")

	// Users
	ErrUserExists   = errors.New("User already exists")
	ErrUserNotFound = errors.New("User not found")

	ErrAccountsDataNotFound = errors.New("Accounts data not found")
	ErrTextDataNotFound     = errors.New("Text data not found")
	ErrBinaryDataNotFound   = errors.New("Binary data not found")
	ErrCardsDataNotFound    = errors.New("Card data not found")

	ErrMetadataNotFound = errors.New("Metadata not found")
)

func TranslateDBError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case pgerrcode.UniqueViolation:
			return ErrConflict
		}
	}
	return err
}
