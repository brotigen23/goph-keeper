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
	time := time.Now()
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
				tableName: accountsTable.tableName,
				rowID:     1,
				data:      "some metadata",
				rows: sqlmock.
					NewRows([]string{metadataTable.idColumnName, metadataTable.createdAtColumnName}).
					AddRow(1, time),
				sqlErr: nil,
			},
			want: want{
				metadata: &model.Metadata{
					ID:        1,
					TableName: accountsTable.tableName,
					RowID:     1,
					Data:      "some metadata",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO "+metadataTable.tableName).
				WithArgs(
					test.args.tableName,
					test.args.rowID,
					test.args.data).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			metadata, err := repo.Create(
				context.Background(),
				test.args.tableName,
				test.args.rowID,
				test.args.data)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.metadata.ID, metadata.ID)
			assert.Equal(t, test.want.metadata.TableName, metadata.TableName)
			assert.Equal(t, test.want.metadata.RowID, metadata.RowID)
			assert.Equal(t, test.want.metadata.Data, metadata.Data)
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
						metadataTable.tableNameColumnName,
						metadataTable.rowIDColumnName,
						metadataTable.dataColumnName,
						metadataTable.createdAtColumnName,
						metadataTable.updatedAtColumnName,
					}).
					AddRow(accountsTable.tableName, 1, "some metadata", time, time),
				sqlErr: nil,
			},
			want: want{
				metadata: &model.Metadata{
					ID:        1,
					TableName: accountsTable.tableName,
					RowID:     1,
					Data:      "some metadata",
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
				metadataTable.tableNameColumnName,
				metadataTable.rowIDColumnName,
				metadataTable.dataColumnName,
				metadataTable.createdAtColumnName,
				metadataTable.updatedAtColumnName,
				metadataTable.tableName,
				metadataTable.idColumnName)

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

func TestGetMetadataByRowID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewMetadataRepository(db, logger.New().Testing())
	type args struct {
		tableName string
		rowID     int
		rows      *sqlmock.Rows
		sqlErr    error
	}
	type want struct {
		metadata model.Metadata
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
				tableName: accountsTable.tableName,
				rowID:     1,
				rows: sqlmock.
					NewRows([]string{
						metadataTable.idColumnName,
						metadataTable.dataColumnName,
						metadataTable.createdAtColumnName,
						metadataTable.updatedAtColumnName,
					}).
					AddRow(1, "some metadata", time, time),
				sqlErr: nil,
			},
			want: want{
				metadata: model.Metadata{
					ID:        1,
					TableName: accountsTable.tableName,
					RowID:     1,
					Data:      "some metadata",
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = \\$1 AND %s = \\$2",
				metadataTable.idColumnName,
				metadataTable.dataColumnName,
				metadataTable.createdAtColumnName,
				metadataTable.updatedAtColumnName,
				metadataTable.tableName,
				metadataTable.tableNameColumnName,
				metadataTable.rowIDColumnName)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.tableName,
					test.args.rowID).
				WillReturnRows(
					test.args.rows)

			metadata, err := repo.GetByRowID(context.Background(), test.args.tableName, test.args.rowID)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.metadata.ID, metadata.ID)
		})
	}
}
