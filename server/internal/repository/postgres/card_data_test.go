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
