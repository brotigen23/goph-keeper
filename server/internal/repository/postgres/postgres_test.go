package postgres

import (
	"context"
	"database/sql/driver"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/stretchr/testify/assert"
)

var timeNow = time.Now()

type args[T model.Model] struct {
	data T
}

type mocks struct {
	// Args with sql query
	args []driver.Value
	// Returned rows from sqlmock
	rows *sqlmock.Rows
}

type want[T model.Model] struct {
	data T
	err  error
}

type testArgs[T model.Model] struct {
	name  string
	mocks mocks
	args  args[T]
	want  want[T]
}

func testPostgresCreate[T model.Model](t *testing.T,
	repository repository.Data[T],
	tests []testArgs[T],
	query string,
	mock sqlmock.Sqlmock) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectQuery(query).
				WithArgs(
					test.mocks.args...).
				WillReturnRows(
					test.mocks.rows)

			mock.ExpectCommit()

			retData, err := repository.Create(
				context.Background(),
				test.args.data,
			)
			assert.Equal(t, test.want.err, err)

			//assert.Equal(t, test.want.data, retData)
			log.Println(retData)
		})
	}
}
