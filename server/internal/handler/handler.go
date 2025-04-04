package handler

import "net/http"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "server response", http.StatusGone)
}
