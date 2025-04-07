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

var timeNow = time.Now()

func TestCreateCardData(t *testing.T) {
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
		card   card
		userID int
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
				card: card{
					number:         "1234",
					cardholderName: "user",
					expire:         timeNow,
					cvv:            "321",
				},
				userID: 1,
				rows: sqlmock.
					NewRows([]string{cardsTable.idColumnName, cardsTable.createdAtColumnName}).
					AddRow(1, timeNow),
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
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO "+cardsTable.tableName).
				WithArgs(
					test.args.userID,
					test.args.card.number,
					test.args.card.cardholderName,
					test.args.card.expire,
					test.args.card.cvv).
				WillReturnRows(
					test.args.rows)

			mock.ExpectCommit()

			card, err := repo.Create(
				context.Background(),
				test.args.userID,
				test.args.card.number,
				test.args.card.cardholderName,
				test.args.card.expire,
				test.want.card.CVV)
			assert.Equal(t, test.want.err, err)

			assert.Equal(t, test.want.card, card)
		})
	}
}

func TestGetCardDataByID(t *testing.T) {
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
				cardsTable.tableName,
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
				cardsTable.tableName,
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
