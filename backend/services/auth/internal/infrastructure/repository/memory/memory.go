package memory

import (
	"context"
	"fmt"

	"github.com/brotigen23/goph-keeper/auth/internal/domain"
)

type repo struct {
	users map[int]domain.User
}

func New() domain.Repository {
	return &repo{
		users: make(map[int]domain.User),
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	for _, v := range r.users {
		if v.Login == user.Login {
			return fmt.Errorf("repo.Create login=%s : %w", user.Login, domain.ErrUserExists)
		}
	}
	user.ID = len(r.users)
	r.users[user.ID] = *user
	return nil
}

func (r *repo) Update(context.Context, domain.Updates) (*domain.User, error) {
	return nil, nil
}

func (r *repo) Get(ctx context.Context, filter domain.Filter) (*domain.User, error) {
	for _, v := range r.users {
		if filter.ID != nil && *filter.ID != v.ID {
			continue
		}
		if filter.Login != nil && *filter.Login != v.Login {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("repo.Get filters = %v : %w", filter, domain.ErrUserNotFound)
}
