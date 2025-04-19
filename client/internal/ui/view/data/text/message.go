package text

import "github.com/brotigen23/goph-keeper/client/internal/domain"

type ServerResponseMsg struct {
	StatusCode int
	Body       string
}

type FetchSuccessMsg struct {
	Accounts []domain.AccountData
}
