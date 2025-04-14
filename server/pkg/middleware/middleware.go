package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/auth"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

type Middleware struct {
	logger *logger.Logger

	accessKey  string
	refreshKey string

	userService *service.UserService
}

func New(logger *logger.Logger, accessKey, refreshKey string) *Middleware {
	return &Middleware{
		logger:     logger,
		accessKey:  accessKey,
		refreshKey: refreshKey,
	}
}

func (m Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusInternalServerError)
			return
		}

		m.logger.Info("Request Body:", "body", string(body))

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)

	})
}

func (m Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, ErrHeaderEmpty.Error(), http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, ErrTokenIsInvalid.Error(), http.StatusUnauthorized)
			return
		}
		// TODO: refresh tokens if valid
		id, err := auth.GetIDFromJWT(token, m.accessKey)
		switch err {
		case nil:
			break
		case auth.ErrTokenIsInvalid:
			http.Error(w, auth.ErrTokenIsInvalid.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(id)
		ctx := context.WithValue(r.Context(), "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
