package memory

import (
	"context"
	"strings"
	"time"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/entity"
	"github.com/brotigen23/goph-keeper/accounts/internal/domain/repo"
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
	account.ID = len(m.items)
	account.CreatedAt = time.Now()
	account.UpdatedAt = account.CreatedAt
	m.items[account.ID] = *account
	return nil
}

func (m *memory) Update(ctx context.Context, updates repo.Updates) (*entity.Account, error) {

	oldItem, ok := m.items[updates.ID]
	if !ok {
		return nil, repo.ErrInvalidID
	}

	newItem := &entity.Account{}

	if updates.Login != nil {
		newItem.Login = *updates.Login
	} else {
		newItem.Login = oldItem.Login
	}

	if updates.Password != nil {
		newItem.Password = *updates.Password
	} else {
		newItem.Password = oldItem.Password
	}

	return newItem, nil
}

func (m *memory) Get(ctx context.Context, filter repo.Filter) ([]entity.Account, error) {
	founded := false
	ret := []entity.Account{}
	for _, v := range m.items {
		if filter.Login != nil && !strings.Contains(v.Login, *filter.Login) {
			continue
		}
		founded = true
		ret = append(ret, v)
	}
	if !founded {
		return nil, repo.ErrNotFound
	}
	return ret, nil
}

func (m *memory) Delete(ctx context.Context, id int) (*entity.Account, error) {
	ret, ok := m.items[id]
	if !ok {
		return nil, repo.ErrNotFound
	}
	delete(m.items, id)
	return &ret, nil
}
