package handler

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/service"
)

type Handler struct {
	service *service.UserDataAggregator

	config *config.Config
}

func New(c *config.Config, a *service.UserDataAggregator) *Handler {
	return &Handler{
		config:  c,
		service: a,
	}
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
