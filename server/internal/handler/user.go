package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/brotigen23/goph-keeper/server/internal/dto"
)

func (h Handler) AccountsDataPost(w http.ResponseWriter, r *http.Request) {}
func (h Handler) AccountsDataGet(w http.ResponseWriter, r *http.Request)  {}

func (h Handler) TextDataPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ErrRequestBodyUnableToRead.Error(), http.StatusInternalServerError)
		return
	}

	textData := &dto.TextData{}
	err = json.Unmarshal(body, textData)
	if err != nil {
		http.Error(w, ErrRequestBodyUnableToRead.Error(), http.StatusInternalServerError)
		return
	}
	// Create

	// Response
}

func (h Handler) TextDataGet(w http.ResponseWriter, r *http.Request) {}

func (h Handler) BinaryData(w http.ResponseWriter, r *http.Request) {}

func (h Handler) CardsData(w http.ResponseWriter, r *http.Request) {}
