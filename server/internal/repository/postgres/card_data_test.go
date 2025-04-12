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

func TestCardCreate(t *testing.T) {
	query := fmt.Sprintf("INSERT INTO %s", cardsTableName)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewCardsRepository(db, logger.New().Testing())

	tests := []testArgs[model.CardData]{
		{
			name: "Test OK",
			mocks: mocks{
				args: []driver.Value{1, "1234", "user", timeNow, "321"},
				rows: sqlmock.
					NewRows([]string{cardsTable.idColumnName, cardsTable.metadataID, cardsTable.createdAtColumnName}).
					AddRow(1, 1, timeNow),
			},
			args: args[model.CardData]{
				data: model.CardData{UserID: 1, Number: "1234", CardholderName: "user", Expire: timeNow, CVV: "321"},
			},
			want: want[model.CardData]{
				data: model.CardData{
					ID:             1,
					UserID:         1,
					Number:         "1234",
					CardholderName: "user",
					Expire:         timeNow,
					CVV:            "321",
					CreatedAt:      timeNow,
					UpdatedAt:      timeNow,
				},
				err: nil,
			},
		},
	}
	testPostgresCreate(t, repo, tests, query, mock)
}

func TestGetCardDataByID(t *testing.T) {
	timeNow := time.Now()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewCardsRepository(db, logger.New().Testing())

	type args struct {
		id     int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		card *model.CardData
		err  error
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
						cardsTable.userIDColumnName,
						cardsTable.numberColumnName,
						cardsTable.cardholderNameColumnName,
						cardsTable.expireColumnName,
						cardsTable.cvvColumnName,
						cardsTable.createdAtColumnName,
						cardsTable.updatedAtColumnName,
					}).
					AddRow(1, "1234", "user", timeNow, "321", timeNow, timeNow),
				sqlErr: nil,
			},
			want: want{
				card: &model.CardData{
					ID:             1,
					UserID:         1,
					Number:         "1234",
					CardholderName: "user",
					Expire:         timeNow,
					CVV:            "321",
					CreatedAt:      timeNow,
					UpdatedAt:      timeNow,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = \\$1",
				cardsTable.userIDColumnName,
				cardsTable.numberColumnName,
				cardsTable.cardholderNameColumnName,
				cardsTable.expireColumnName,
				cardsTable.cvvColumnName,
				cardsTable.createdAtColumnName,
				cardsTable.updatedAtColumnName,
				cardsTableName,
				cardsTable.idColumnName)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.id).
				WillReturnRows(
					test.args.rows)

			textData, err := repo.GetByID(context.Background(), test.args.id)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.card, textData)
		})
	}
}

func TestGetCardDataByUserID(t *testing.T) {
	timeNow := time.Now()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	var repo = NewCardsRepository(db, logger.New().Testing())
	type card struct {
		number         string
		cardholderName string
		expire         time.Time
		cvv            string
	}

	type args struct {
		userID int
		rows   *sqlmock.Rows
		sqlErr error
	}
	type want struct {
		cards []model.CardData
		err   error
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
						cardsTable.idColumnName,
						cardsTable.numberColumnName,
						cardsTable.cardholderNameColumnName,
						cardsTable.expireColumnName,
						cardsTable.cvvColumnName,
						cardsTable.createdAtColumnName,
						cardsTable.updatedAtColumnName,
					}).
					AddRow(1, "1234", "user", timeNow, "321", timeNow, timeNow).
					AddRow(2, "5678", "user", timeNow, "213", timeNow, timeNow).
					AddRow(3, "9101", "user", timeNow, "123", timeNow, timeNow),
				sqlErr: nil,
			},
			want: want{
				cards: []model.CardData{
					{ID: 1, UserID: 1, Number: "1234", CardholderName: "user",
						Expire: timeNow, CVV: "321", CreatedAt: timeNow, UpdatedAt: timeNow},
					{ID: 2, UserID: 1, Number: "5678", CardholderName: "user",
						Expire: timeNow, CVV: "213", CreatedAt: timeNow, UpdatedAt: timeNow},
					{ID: 3, UserID: 1, Number: "9101", CardholderName: "user",
						Expire: timeNow, CVV: "123", CreatedAt: timeNow, UpdatedAt: timeNow},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = \\$1",
				cardsTable.idColumnName,
				cardsTable.numberColumnName,
				cardsTable.cardholderNameColumnName,
				cardsTable.expireColumnName,
				cardsTable.cvvColumnName,
				cardsTable.createdAtColumnName,
				cardsTable.updatedAtColumnName,
				cardsTableName,
				cardsTable.userIDColumnName)

			mock.ExpectQuery(query).
				WithArgs(
					test.args.userID).
				WillReturnRows(
					test.args.rows)

			cards, err := repo.GetByUserID(context.Background(), test.args.userID)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.cards, cards)
		})
	}
}
