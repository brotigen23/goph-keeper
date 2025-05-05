package database

import (
	"github.com/brotigen23/goph-keeper/server/pkg/migration"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

func New(driver, dsn string) (*Database, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (d *Database) Migrate(path string) error {
	err := migration.Migrate(d.DB.DB, path)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DiMigrate(path string) error {
	err := migration.Migrate(d.DB.DB, path)
	if err != nil {
		return err
	}
	return nil
}

func (d Database) Close() {
	d.DB.Close()
}
