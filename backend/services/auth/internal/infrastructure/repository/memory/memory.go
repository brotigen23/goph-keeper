package memory

import (
	"context"

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
			return domain.ErrUserExists
		}
	}
	user.ID = len(r.users)
	r.users[user.ID] = *user
	return nil
}

func (r *repo) Update(context.Context, domain.Updates) (*domain.User, error) {
	return nil, nil
}

func (r *repo) Get(context.Context, domain.Filter) (*domain.User, error) {
	return nil, nil
}
