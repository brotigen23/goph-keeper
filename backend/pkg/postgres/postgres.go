package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Conn    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
func New(DSN string) (*Postgres, error) {
	pgxp, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		Conn:    pgxp,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}
