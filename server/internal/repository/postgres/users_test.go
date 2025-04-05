package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/crypt"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo repository.Users = NewUsers(db, nil)
	type userCredentials struct {
		login    string
		password string
	}

	type args struct {
		userCredentials userCredentials
		rows            *sqlmock.Rows
		pqErr           *pq.Error
	}
	type want struct {
		user *model.User
		err  error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				userCredentials: userCredentials{
					login:    "user1",
					password: "pass1",
				},
				rows: sqlmock.
					NewRows([]string{"id", "created_at"}).
					AddRow(1, time),
				pqErr: nil,
			},
			want: want{
				user: &model.User{
					ID:    1,
					Login: "user1",
				},
			},
		},
		{
			name: "Test Conflict",
			args: args{
				pqErr: &pq.Error{
					Code: pgerrcode.UniqueViolation,
				},
			},
			want: want{
				user: &model.User{
					ID:    1,
					Login: "user1",
				},
				err: repository.ErrUserExists,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.args.pqErr {
			case nil:
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(
						test.args.userCredentials.login,
						sqlmock.AnyArg()).
					WillReturnRows(
						test.args.rows)

				mock.ExpectCommit()
				user, err := repo.Create(
					context.Background(),
					test.args.userCredentials.login,
					test.args.userCredentials.password)
				assert.Equal(t, test.want.err, err)

				assert.Equal(t, test.want.user.ID, user.ID)
				assert.Equal(t, test.want.user.Login, user.Login)
				assert.NoError(t, crypt.CheckPasswordHash(test.args.userCredentials.password, user.Password))
			default:
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(
						test.args.userCredentials.login,
						sqlmock.AnyArg()).
					WillReturnError(test.args.pqErr)

				mock.ExpectRollback()

				_, err := repo.Create(
					context.Background(),
					test.args.userCredentials.login,
					test.args.userCredentials.password)

				assert.Equal(t, test.want.err, err)
			}
		})
	}
}
