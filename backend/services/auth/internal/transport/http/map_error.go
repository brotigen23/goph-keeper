package http

import (
	"errors"
	"net/http"

	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/response"
	"github.com/brotigen23/goph-keeper/auth/internal/usecase"
)

func MapError(err error) (int, any) {
	switch {
	case errors.Is(err, usecase.ErrEmptyLogin):
		return http.StatusBadRequest, &response.Err{Msg: "login must be not empty"}
	case errors.Is(err, usecase.ErrEmptyPassword):
		return http.StatusBadRequest, &response.Err{Msg: "password must be not empty"}
	case errors.Is(err, usecase.ErrUserNotFound):
		return http.StatusBadRequest, &response.Err{Msg: "user not found"}

	default:
		return http.StatusInternalServerError, &response.Err{Msg: "internal error"}
	}
}
