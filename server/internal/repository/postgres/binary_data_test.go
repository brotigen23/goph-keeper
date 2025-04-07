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

func TestCreateBinaryData(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewBinaryRepository(db, logger.New().Testing())
	type args struct {
		userID int
		data   []byte
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		binaryData *model.BinaryData
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
				data:   []byte("some data"),
				rows: sqlmock.
					NewRows([]string{binaryTable.idColumnName, binaryTable.createdAtColumnName}).
					AddRow(1, time),
				sqlErr: nil,
			},
			want: want{
				binaryData: &model.BinaryData{
					ID:     1,
					UserID: 1,
					Data:   []byte("some data"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO "+binaryTable.tableName).
				WithArgs(
					test.args.userID,
					test.args.data).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			binaryData, err := repo.Create(
				context.Background(),
				test.args.userID,
				test.args.data)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.binaryData.ID, binaryData.ID)
			assert.Equal(t, test.want.binaryData.UserID, binaryData.UserID)
			assert.Equal(t, test.want.binaryData.Data, binaryData.Data)
		})
	}
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
						binaryTable.userIDColumnName,
						binaryTable.dataColumnName,
						binaryTable.createdAtColumnName,
						binaryTable.updatedAtColumnName,
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
				binaryTable.userIDColumnName,
				binaryTable.dataColumnName,
				binaryTable.createdAtColumnName,
				binaryTable.updatedAtColumnName,
				binaryTable.tableName,
				binaryTable.idColumnName)

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
						binaryTable.idColumnName,
						binaryTable.dataColumnName,
						binaryTable.createdAtColumnName,
						binaryTable.updatedAtColumnName,
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
				binaryTable.idColumnName,
				binaryTable.dataColumnName,
				binaryTable.createdAtColumnName,
				binaryTable.updatedAtColumnName,
				binaryTable.tableName,
				binaryTable.userIDColumnName)

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
