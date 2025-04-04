package migration

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(path, "pq", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	errSource, errDB := m.Close()
	if errSource != nil {
		return errSource
	}
	if errDB != nil {
		return errDB
	}
	return nil
}
