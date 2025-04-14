package postgres

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
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
