package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/mapper"
)

func (h Handler) MetadataPut(w http.ResponseWriter, r *http.Request) {
	newMetadata := &dto.MetadataPut{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newMetadata)
	if err != nil {
		http.Error(w, ErrRequestBodyUnableToRead.Error(), http.StatusBadRequest)
		return
	}
	if newMetadata.Data == nil {
		http.Error(w, "Nothing to update", http.StatusBadRequest)
		return
	}
	savedMetadata, err := h.service.UpdateMetadata(context.Background(),
		newMetadata.ID, *newMetadata.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto := mapper.MetadataToDTO(*savedMetadata)

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}
