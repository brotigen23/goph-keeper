package handler

import "net/http"

func (h Handler) AccountsDataPost(w http.ResponseWriter, r *http.Request) {}
func (h Handler) AccountsDataGet(w http.ResponseWriter, r *http.Request)  {}

func (h Handler) TextDataPost(w http.ResponseWriter, r *http.Request) {}
func (h Handler) TextDataGet(w http.ResponseWriter, r *http.Request)  {}

func (h Handler) BinaryData(w http.ResponseWriter, r *http.Request) {}

func (h Handler) CardsData(w http.ResponseWriter, r *http.Request) {}
