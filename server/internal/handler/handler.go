package handler

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/service"
)

type Handler struct {
	userService *service.UserService
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetUserService(service *service.UserService) {
	h.userService = service
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "server response", http.StatusGone)
}
