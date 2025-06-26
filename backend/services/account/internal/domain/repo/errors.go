package repo

import "errors"

var (
	ErrInvalidInput = errors.New("")
	ErrInvalidID    = errors.New("record doesnt exist")
	ErrNotFound     = errors.New("not found any rows")
)
