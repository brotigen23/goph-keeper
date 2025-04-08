package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/auth"
	"github.com/brotigen23/goph-keeper/server/pkg/response"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &dto.UserDTO{}
	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	savedUser, err := h.service.CreateNewUser(context.Background(), user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessExpire := time.Hour * 24
	refreshExpire := time.Hour * 72

	accessToken, refreshToken, err := auth.CreateTokens(
		savedUser.ID,
		h.config.JWT.AccessKey,
		h.config.JWT.RefreshKey,
		accessExpire, refreshExpire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.CreateTokenHeaders(w, accessToken, refreshToken)
	w.WriteHeader(http.StatusAccepted)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &dto.UserDTO{}
	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gotUser, err := h.service.ValidateUserLogin(context.Background(), user.Login, user.Password)
	switch err {
	case nil:
		break
	case service.ErrUserNotFound:
		http.Error(w, ErrUserNotFound.Error(), http.StatusUnauthorized)
		return
	case service.ErrIncorrectPassword:
		http.Error(w, ErrIncorrectPassword.Error(), http.StatusUnauthorized)
		return
	default:

	}

	accessExpire := time.Hour * 24
	refreshExpire := time.Hour * 72

	accessToken, refreshToken, err := auth.CreateTokens(
		gotUser.ID,
		h.config.JWT.AccessKey,
		h.config.JWT.RefreshKey,
		accessExpire, refreshExpire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.CreateTokenHeaders(w, accessToken, refreshToken)
	w.WriteHeader(http.StatusAccepted)
}
