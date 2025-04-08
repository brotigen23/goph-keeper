package database

import (
	"database/sql"

	"github.com/brotigen23/goph-keeper/server/pkg/migration"
)

type Database struct {
	DB *sql.DB
}

func New(driver, dsn string) (*Database, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (d *Database) Migrate(path string) error {
	err := migration.Migrate(d.DB, path)
	if err != nil {
		return err
	}
	return nil
}

func (d Database) Close() {
	d.DB.Close()
}
