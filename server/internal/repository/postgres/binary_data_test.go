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

func TestBinaryCreate(t *testing.T) {
	query := fmt.Sprintf("INSERT INTO %s", binaryTableName)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewBinaryRepository(db, logger.New().Testing())

	tests := []testArgs[model.BinaryData]{
		{
			name: "Test OK",
			mocks: mocks{
				args: []driver.Value{1, []byte("some data")},
				rows: sqlmock.
					NewRows([]string{binaryTable.id, binaryTable.metadataID, binaryTable.createdAt}).
					AddRow(1, 1, timeNow),
			},
			args: args[model.BinaryData]{
				data: model.BinaryData{UserID: 1, Data: []byte("some data")},
			},
			want: want[model.BinaryData]{
				data: model.BinaryData{
					ID:         1,
					UserID:     1,
					MetadataID: 1,
					Data:       []byte("some data"),
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
				},
				err: nil,
			},
		},
	}
	testPostgresCreate(t, repo, tests, query, mock)
}
