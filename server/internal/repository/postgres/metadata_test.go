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

func TestCreateMetadata(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewMetadataRepository(db, logger.New().Testing())
	type args struct {
		tableName string
		rowID     int
		data      string
		rows      *sqlmock.Rows
		sqlErr    error
	}
	type want struct {
		metadata *model.Metadata
		err      error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				tableName: accountsTableName,
				rowID:     1,
				data:      "some metadata",
				rows: sqlmock.
					NewRows([]string{metadataTable.id, metadataTable.createdAt}).
					AddRow(1, timeNow),
				sqlErr: nil,
			},
			want: want{
				metadata: &model.Metadata{
					ID:        1,
					Data:      "some metadata",
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO " + metadataTableName).
				WithArgs(
					test.args.data).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			metadata, err := repo.Create(
				context.Background(),
				test.args.data)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.metadata, metadata)
		})
	}
}

func TestGetMetadataByID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewMetadataRepository(db, logger.New().Testing())
	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		metadata *model.Metadata
		err      error
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
						metadataTable.data,
						metadataTable.createdAt,
						metadataTable.updatedAt,
					}).
					AddRow("some metadata", time, time),
				sqlErr: nil,
			},
			want: want{
				metadata: &model.Metadata{
					ID:        1,
					Data:      "some metadata",
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = \\$1",
				metadataTable.data,
				metadataTable.createdAt,
				metadataTable.updatedAt,
				metadataTableName,
				metadataTable.id)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.id).
				WillReturnRows(
					test.args.rows)

			textData, err := repo.GetByID(context.Background(), test.args.id)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.metadata, textData)
		})
	}
}
