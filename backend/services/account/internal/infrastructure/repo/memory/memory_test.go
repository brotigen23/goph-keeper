package memory

import (
	"context"
	"testing"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/entity"
	"github.com/brotigen23/goph-keeper/accounts/internal/domain/repo"
	"github.com/brotigen23/goph-keeper/shared/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		*entity.Account
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Test OK",
			args: args{
				&entity.Account{
					Login:    "User1",
					Password: "Pass1",
				},
			},
			err: nil,
		},
	}

	repo := NewMemory()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Create(context.Background(), tt.args.Account)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	type args struct {
		repo.Updates
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Test OK",
			args: args{
				repo.Updates{
					ID:       0,
					Login:    util.StrPtr("UUU"),
					Password: util.StrPtr("PPP"),
				},
			},
			err: nil,
		},
		{
			name: "Test Invalid ID",
			args: args{
				repo.Updates{
					ID: 111,
				},
			},
			err: repo.ErrInvalidID,
		},
	}

	repo := NewMemory()
	err := repo.Create(context.Background(), &entity.Account{Login: "User1", Password: "Pass2"})
	assert.NoError(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := repo.Update(context.Background(), tt.args.Updates)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.Login, account.Login)
			}
		})
	}
}

func TestGet(t *testing.T) {

	type args struct {
		repo.Filter
	}
	type want struct {
		length int
		err    error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test No Rows Founded",
			args: args{
				repo.Filter{
					Login: util.StrPtr("P"),
				},
			},
			want: want{
				length: 0,
				err:    repo.ErrNotFound,
			},
		},
		{
			name: "Test OK all rows",
			want: want{
				length: 3,
				err:    nil,
			},
		},
		{
			name: "Test OK all rows with 'user'",
			args: args{
				repo.Filter{
					Login: util.StrPtr("User"),
				},
			},
			want: want{
				length: 3,
				err:    nil,
			},
		},
	}

	repo := NewMemory()
	err := repo.Create(context.Background(), &entity.Account{Login: "User1", Password: "Pass2"})
	assert.NoError(t, err)
	err = repo.Create(context.Background(), &entity.Account{Login: "User2", Password: "Pass2"})
	assert.NoError(t, err)
	err = repo.Create(context.Background(), &entity.Account{Login: "User3", Password: "Pass2"})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Get(context.Background(), tt.args.Filter)
			if tt.want.err != nil {
				assert.ErrorIs(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)

			}
		})
	}
}

func TestDelete(t *testing.T) {

	type args struct {
		ID int
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Test OK",
			args: args{
				ID: 1,
			},
			err: nil,
		},
		{
			name: "Test Not Exist",
			args: args{
				ID: 111,
			},
			err: repo.ErrNotFound,
		},
	}

	repo := &memory{items: map[int]entity.Account{}}
	repo.items[1] = entity.Account{ID: 1, Login: "Login", Password: "Password"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Delete(context.Background(), tt.args.ID)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
