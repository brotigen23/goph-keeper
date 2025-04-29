package ui

import "github.com/brotigen23/goph-keeper/client/internal/app/domain"

type FetchSuccessMsg[T domain.Model] struct {
	Data []T
}
