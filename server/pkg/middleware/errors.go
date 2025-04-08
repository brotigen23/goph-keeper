package middleware

import "errors"

var (
	ErrHeaderEmpty = errors.New("Header is empty")

	ErrTokenIsInvalid = errors.New("JWT is empty")
)
