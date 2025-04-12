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

func TestGetBinaryDataByID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewBinaryRepository(db, logger.New().Testing())
	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		textData *model.BinaryData
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
						binaryTable.userID,
						binaryTable.data,
						binaryTable.createdAt,
						binaryTable.updatedAt,
					}).
					AddRow(1, []byte("data1"), time, time),
				sqlErr: nil,
			},
			want: want{
				textData: &model.BinaryData{
					ID:        1,
					UserID:    1,
					Data:      []byte("data1"),
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
				binaryTable.userID,
				binaryTable.data,
				binaryTable.createdAt,
				binaryTable.updatedAt,
				binaryTableName,
				binaryTable.id)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.id).
				WillReturnRows(
					test.args.rows)

			binaryData, err := repo.GetByID(context.Background(), test.args.id)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.textData, binaryData)
		})
	}

}
func TestGetBinaryDataByUserID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewBinaryRepository(db, logger.New().Testing())
	type args struct {
		userID int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		binaryData []model.BinaryData
		err        error
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
						binaryTable.id,
						binaryTable.data,
						binaryTable.createdAt,
						binaryTable.updatedAt,
					}).
					AddRow(1, []byte("data1"), time, time).
					AddRow(2, []byte("data2"), time, time),
				sqlErr: nil,
			},
			want: want{
				binaryData: []model.BinaryData{
					{
						ID:        1,
						UserID:    1,
						Data:      []byte("data1"),
						CreatedAt: time,
						UpdatedAt: time,
					},
					{
						ID:        2,
						UserID:    1,
						Data:      []byte("data2"),
						CreatedAt: time,
						UpdatedAt: time,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = \\$1",
				binaryTable.id,
				binaryTable.data,
				binaryTable.createdAt,
				binaryTable.updatedAt,
				binaryTableName,
				binaryTable.userID)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.userID).
				WillReturnRows(
					test.args.rows)

			binaryData, err := repo.GetByUserID(context.Background(), test.args.userID)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.binaryData, binaryData)
		})
	}
}
