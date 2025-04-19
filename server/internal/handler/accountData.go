package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/mapper"
	"github.com/brotigen23/goph-keeper/server/internal/service"
)

func (h Handler) AccountsDataPost(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, ErrIncorrectUserID.Error(), http.StatusBadRequest)
		return
	}
	data := &dto.AccountsGet{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, ErrRequestBodyUnableToRead.Error(), http.StatusBadRequest)
		return
	}
	model, err := h.service.CreateAccount(context.Background(), userID, data.Login, data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto := mapper.AccountToDTO(model)

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

func (h Handler) AccountsDataGet(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, ErrIncorrectUserID.Error(), http.StatusBadRequest)
		return
	}
	model, metadata, err := h.service.GetAccounts(context.Background(), userID)
	switch err {
	case nil:
		break
	case service.ErrDataNotFound:
		http.Error(w, "No content", http.StatusNoContent)
		return
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto := mapper.AccountsToDTO(model, metadata)
	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusAccepted)
}

func (h Handler) AccountsDataPut(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, ErrIncorrectUserID.Error(), http.StatusBadRequest)
		return
	}
	newAccount := &dto.AccountPut{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newAccount)
	if err != nil {
		http.Error(w, ErrRequestBodyUnableToRead.Error(), http.StatusBadRequest)
		return
	}
	savedAccount, err := h.service.UpdateAccount(context.Background(),
		userID, newAccount.ID, newAccount.Login, newAccount.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto := mapper.AccountPutToDTO(savedAccount)

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}
