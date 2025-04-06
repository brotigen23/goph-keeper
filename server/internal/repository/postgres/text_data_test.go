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

func TestCreateTextData(t *testing.T) {
	time := time.Now()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewTextDataRepository(db, logger.New().Testing())
	type args struct {
		userID int
		data   string
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
				userID: 1,
				data:   "some data",
				rows: sqlmock.
					NewRows([]string{textDataTable.idColumnName, textDataTable.createdAtColumnName}).
					AddRow(1, time),
				sqlErr: nil,
			},
			want: want{
				textData: &model.TextData{
					ID:     1,
					UserID: 1,
					Data:   "some data",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO text_data").
				WithArgs(
					test.args.userID,
					test.args.data).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			user, err := repo.Create(
				context.Background(),
				test.args.userID,
				test.args.data)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.textData.ID, user.ID)
			assert.Equal(t, test.want.textData.UserID, user.UserID)
			assert.Equal(t, test.want.textData.Data, user.Data)
		})
	}
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
						textDataTable.userIDColumnName,
						textDataTable.dataColumnName,
						textDataTable.createdAtColumnName,
						textDataTable.updatedAtColumnName,
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
				textDataTable.userIDColumnName,
				textDataTable.dataColumnName,
				textDataTable.createdAtColumnName,
				textDataTable.updatedAtColumnName,
				textDataTable.tableName,
				textDataTable.idColumnName)

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
						textDataTable.idColumnName,
						textDataTable.dataColumnName,
						textDataTable.createdAtColumnName,
						textDataTable.updatedAtColumnName,
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
				textDataTable.idColumnName,
				textDataTable.dataColumnName,
				textDataTable.createdAtColumnName,
				textDataTable.updatedAtColumnName,
				textDataTable.tableName,
				textDataTable.userIDColumnName)

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
