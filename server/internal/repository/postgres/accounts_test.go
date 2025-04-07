package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountsData(t *testing.T) {
	timeNow := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewAccountsRepository(db, logger.New().Testing())

	type accountData struct {
		login    string
		password string
	}
	type args struct {
		userID      int
		accountData accountData
		rows        *sqlmock.Rows
		sqlErr      error
	}
	type want struct {
		accountData *model.AccountData
		err         error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				userID:      1,
				accountData: accountData{login: "user1", password: "pass1"},
				rows: sqlmock.
					NewRows([]string{accountsTable.idColumnName, accountsTable.createdAtColumnName}).
					AddRow(1, timeNow),
				sqlErr: nil,
			},
			want: want{
				accountData: &model.AccountData{
					ID:        1,
					UserID:    1,
					Login:     "user1",
					Password:  "pass1",
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO "+accountsTable.tableName).
				WithArgs(
					test.args.userID,
					test.args.accountData.login,
					test.args.accountData.password).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			accountData, err := repo.Create(
				context.Background(),
				test.args.userID,
				test.args.accountData.login,
				test.args.accountData.password)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.accountData, accountData)
		})
	}
}

func TestGetAccountDataByID(t *testing.T) {
	timeNow := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewAccountsRepository(db, logger.New().Testing())
	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		accountData *model.AccountData
		err         error
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
						accountsTable.userIDColumnName,
						accountsTable.loginColumnName,
						accountsTable.passwordColumnName,
						accountsTable.createdAtColumnName,
						accountsTable.updatedAtColumnName,
					}).
					AddRow(1, "user1", "pass1", timeNow, timeNow),
				sqlErr: nil,
			},
			want: want{
				accountData: &model.AccountData{
					ID:        1,
					UserID:    1,
					Login:     "user1",
					Password:  "pass1",
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = \\$1",
				accountsTable.userIDColumnName,
				accountsTable.loginColumnName,
				accountsTable.passwordColumnName,
				accountsTable.createdAtColumnName,
				accountsTable.updatedAtColumnName,
				accountsTable.tableName,
				accountsTable.idColumnName)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.id).
				WillReturnRows(
					test.args.rows)

			accountData, err := repo.GetByID(context.Background(), test.args.id)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.accountData, accountData)
		})
	}
}

func TestGetAccountsDataByUserID(t *testing.T) {
	timeNow := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewAccountsRepository(db, logger.New().Testing())
	type args struct {
		userID int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		accountsData []model.AccountData
		err          error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				userID: 1,
				rows: sqlmock.
					NewRows([]string{
						accountsTable.idColumnName,
						accountsTable.loginColumnName,
						accountsTable.passwordColumnName,
						accountsTable.createdAtColumnName,
						accountsTable.updatedAtColumnName,
					}).
					AddRow(1, "user1", "pass1", timeNow, timeNow).
					AddRow(2, "user2", "pass2", timeNow, timeNow),
				sqlErr: nil,
			},
			want: want{
				accountsData: []model.AccountData{
					{
						ID:        1,
						UserID:    1,
						Login:     "user1",
						Password:  "pass1",
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
					{
						ID:        2,
						UserID:    1,
						Login:     "user2",
						Password:  "pass2",
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = \\$1",
				accountsTable.idColumnName,
				accountsTable.loginColumnName,
				accountsTable.passwordColumnName,
				accountsTable.createdAtColumnName,
				accountsTable.updatedAtColumnName,
				accountsTable.tableName,
				accountsTable.userIDColumnName)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.userID).
				WillReturnRows(
					test.args.rows)

			accountsData, err := repo.GetByUserID(context.Background(), test.args.userID)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.accountsData, accountsData)
		})
	}
}
