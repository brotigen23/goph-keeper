package account

import "github.com/brotigen23/goph-keeper/server/internal/dto"

type PostRequest struct {
	Model
}

type PostResponse struct {
	dto.BaseData
	Model
}
