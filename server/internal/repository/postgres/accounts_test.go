package postgres

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSome(t *testing.T) {
	query := fmt.Sprintf("INSERT INTO %s", accountsTableName)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewAccountsRepository(db, logger.New().Testing())

	tests := []testArgs[model.AccountData]{
		{
			name: "Test OK",
			mocks: mocks{
				args: []driver.Value{1, "user1", "pass1"},
				rows: sqlmock.
					NewRows([]string{accountsTable.id, accountsTable.metadataID, accountsTable.createdAt}).
					AddRow(1, 1, timeNow),
			},
			args: args[model.AccountData]{
				data: model.AccountData{UserID: 1, Login: "user1", Password: "pass1"},
			},
			want: want[model.AccountData]{
				data: model.AccountData{
					ID:         1,
					UserID:     1,
					MetadataID: 1,
					Login:      "user1",
					Password:   "pass1",
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
				},
				err: nil,
			},
		},
	}
	testPostgresCreate(t, repo, tests, query, mock)
}

func TestCreateAccountsData(t *testing.T) {
	timeNow := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewAccountsRepository(db, logger.New().Testing())

	type args struct {
		accountData model.AccountData
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
				accountData: model.AccountData{UserID: 1, Login: "user1", Password: "pass1"},
				rows: sqlmock.
					NewRows([]string{accountsTable.id, accountsTable.createdAt}).
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
			mock.ExpectQuery("INSERT INTO "+accountsTableName).
				WithArgs(
					test.args.accountData.UserID,
					test.args.accountData.Login,
					test.args.accountData.Password).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			accountData, err := repo.Create(
				context.Background(),
				test.args.accountData,
			)
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
						accountsTable.userID,
						accountsTable.login,
						accountsTable.password,
						accountsTable.createdAt,
						accountsTable.updatedAt,
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
				accountsTable.userID,
				accountsTable.login,
				accountsTable.password,
				accountsTable.createdAt,
				accountsTable.updatedAt,
				accountsTableName,
				accountsTable.id)

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
						accountsTable.id,
						accountsTable.login,
						accountsTable.password,
						accountsTable.createdAt,
						accountsTable.updatedAt,
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
				accountsTable.id,
				accountsTable.login,
				accountsTable.password,
				accountsTable.createdAt,
				accountsTable.updatedAt,
				accountsTableName,
				accountsTable.userID)

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
