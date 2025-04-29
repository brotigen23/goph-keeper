package datacontroller

import "github.com/brotigen23/goph-keeper/client/internal/app/domain"

type CRUDMsg[T domain.Model] struct {
	Action int
	Item   T
}

type RequestDataMsg[T domain.Model] struct{}
