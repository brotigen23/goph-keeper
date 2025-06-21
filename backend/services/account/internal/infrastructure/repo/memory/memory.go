package memory

import (
	"context"
	"time"

	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/entity"
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/repo"
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/infrastructure/repo/dto"
)

type memory struct {
	items map[int]entity.Account
}

func NewMemory() repo.Repository {
	return &memory{
		items: make(map[int]entity.Account),
	}
}

func (m *memory) Create(ctx context.Context, account *entity.Account) error {
	a := dto.NewAccountFromEntity(*account)
	a.ID = len(m.items)
	a.CreatedAt = time.Now()
	a.UpdatedAt = a.CreatedAt
	m.items[a.ID] = *a.ToEntity()
	return nil
}

func (m *memory) Get(context.Context, repo.Filter) ([]entity.Account, error) {
	return nil, nil
}
