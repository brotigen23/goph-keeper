package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
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

	var repo repository.Users = NewUsers(db, logger.New().Testing())
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
					ID:       1,
					Login:    "user1",
					Password: "pass1",
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
				assert.Equal(t, test.want.user.Password, user.Password)
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
func TestGetByID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo repository.Users = NewUsers(db, logger.New().Testing())
	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
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
				id: 1,
				rows: sqlmock.
					NewRows([]string{
						userTable.idColumnName,
						userTable.loginColumnName,
						userTable.passwordColumnName,
						userTable.createdAtColumnName,
						userTable.updatedAtColumnName}).
					AddRow(
						1,
						"user1",
						"pass1",
						time,
						time,
					),
				sqlErr: nil,
			},
			want: want{
				user: &model.User{
					ID:        1,
					Login:     "user1",
					Password:  "pass1",
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
		{
			name: "Test Not Found",
			args: args{
				id:     2,
				sqlErr: sql.ErrNoRows,
			},
			want: want{
				err: repository.ErrUserNotFound,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
				userTable.idColumnName,
				userTable.loginColumnName,
				userTable.passwordColumnName,
				userTable.createdAtColumnName,
				userTable.updatedAtColumnName,
				userTable.tableName,
				userTable.idColumnName)

			switch test.args.sqlErr {
			case nil:
				mock.ExpectQuery(query).
					WithArgs(
						test.args.id).
					WillReturnRows(
						test.args.rows)

				user, err := repo.GetByID(context.Background(), test.args.id)
				assert.Equal(t, test.want.err, err)

				assert.Equal(t, test.want.user, user)
			default:
				mock.ExpectQuery(query).
					WithArgs(
						test.args.id).
					WillReturnError(
						test.args.sqlErr)

				user, err := repo.GetByID(context.Background(), test.args.id)
				assert.Equal(t, test.want.err, err)

				assert.Equal(t, test.want.user, user)
			}
		})
	}
}
