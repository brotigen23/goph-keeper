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

func TestBinarysCreate(t *testing.T) {
	query := fmt.Sprintf("INSERT INTO %s", textTableName)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewTextDataRepository(db, logger.New().Testing())

	tests := []testArgs[model.TextData]{
		{
			name: "Test OK",
			mocks: mocks{
				args: []driver.Value{1, "some data"},
				rows: sqlmock.
					NewRows([]string{textDataTable.id, textDataTable.metadataID, textDataTable.createdAt}).
					AddRow(1, 1, timeNow),
			},
			args: args[model.TextData]{
				data: model.TextData{UserID: 1, Data: "some data"},
			},
			want: want[model.TextData]{
				data: model.TextData{
					ID:         1,
					UserID:     1,
					MetadataID: 1,
					Data:       "some data",
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
				},
				err: nil,
			},
		},
	}
	testPostgresCreate(t, repo, tests, query, mock)
}
