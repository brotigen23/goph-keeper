package manager

import "github.com/brotigen23/goph-keeper/client/internal/core/domain"

type FetchSuccessMsg[T domain.Model] struct {
	Data []T
}
