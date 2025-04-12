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

func TestGetTextDataByID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewTextDataRepository(db, logger.New().Testing())
	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		textData *model.TextData
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
						textDataTable.userID,
						textDataTable.data,
						textDataTable.createdAt,
						textDataTable.updatedAt,
					}).
					AddRow(1, "data1", time, time),
				sqlErr: nil,
			},
			want: want{
				textData: &model.TextData{
					ID:        1,
					UserID:    1,
					Data:      "data1",
					CreatedAt: time,
					UpdatedAt: time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
				textDataTable.userID,
				textDataTable.data,
				textDataTable.createdAt,
				textDataTable.updatedAt,
				textTableName,
				textDataTable.id)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.id).
				WillReturnRows(
					test.args.rows)

			textData, err := repo.GetByID(context.Background(), test.args.id)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.textData, textData)
		})
	}

}
func TestGetTextDataByUserID(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewTextDataRepository(db, logger.New().Testing())
	type args struct {
		userID int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		textData []model.TextData
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
				userID: 1,
				rows: sqlmock.
					NewRows([]string{
						textDataTable.id,
						textDataTable.data,
						textDataTable.createdAt,
						textDataTable.updatedAt,
					}).
					AddRow(1, "data1", time, time).
					AddRow(2, "data2", time, time),
				sqlErr: nil,
			},
			want: want{
				textData: []model.TextData{
					{
						ID:        1,
						UserID:    1,
						Data:      "data1",
						CreatedAt: time,
						UpdatedAt: time,
					},
					{
						ID:        2,
						UserID:    1,
						Data:      "data2",
						CreatedAt: time,
						UpdatedAt: time,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
				textDataTable.id,
				textDataTable.data,
				textDataTable.createdAt,
				textDataTable.updatedAt,
				textTableName,
				textDataTable.userID)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.userID).
				WillReturnRows(
					test.args.rows)

			textData, err := repo.GetByUserID(context.Background(), test.args.userID)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.textData, textData)
		})
	}
}
