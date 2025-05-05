package account

import (
	"context"
	"fmt"
	"log"

	"github.com/brotigen23/goph-keeper/server/db/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/jmoiron/sqlx"
)

// TODO: Implement all stuff

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.AccountsData {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, item *model.Account) error {
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s) VALUES(:%s, :%s, :%s, :%s) RETURNING %s, %s",
		postgres.AccountsTable.Name,

		postgres.AccountsTable.Columns.Login,
		postgres.AccountsTable.Columns.Password,
		postgres.AccountsTable.Columns.UserID,
		postgres.AccountsTable.Columns.Metadata,

		postgres.AccountsTable.Columns.Login,
		postgres.AccountsTable.Columns.Password,
		postgres.AccountsTable.Columns.UserID,
		postgres.AccountsTable.Columns.Metadata,

		postgres.AccountsTable.Columns.ID,
		postgres.AccountsTable.Columns.CreatedAt)
	rows, err := r.db.NamedQueryContext(ctx, query, item)
	if err != nil {
		return repository.TranslateDBError(err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(item); err != nil {
			return repository.TranslateDBError(err)
		}
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, id int) (*model.Account, error) {
	return nil, repository.ErrNotImplement
}

func (r *Repository) Update(ctx context.Context, item *model.Account) error {
	// Igrones created_at, updated_at and user_id

	query := fmt.Sprintf("UPDATE %s SET %s = :%s, %s = :%s, %s = :%s WHERE %s = :%s",
		postgres.AccountsTable.Name,
		postgres.AccountsTable.Columns.Login,
		postgres.AccountsTable.Columns.Login,
		postgres.AccountsTable.Columns.Password,
		postgres.AccountsTable.Columns.Password,
		postgres.AccountsTable.Columns.Metadata,
		postgres.AccountsTable.Columns.Metadata,
		postgres.AccountsTable.Columns.ID,
		postgres.AccountsTable.Columns.ID,
	)
	_, err := r.db.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return repository.ErrNotImplement
}

func (r *Repository) GetAll(ctx context.Context, userID int) ([]model.Account, error) {
	ret := []model.Account{}
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		postgres.AccountsTable.Columns.ID,
		postgres.AccountsTable.Columns.Login,
		postgres.AccountsTable.Columns.Password,
		postgres.AccountsTable.Columns.UserID,
		postgres.AccountsTable.Columns.Metadata,
		postgres.AccountsTable.Columns.CreatedAt,
		postgres.AccountsTable.Columns.UpdatedAt,

		postgres.AccountsTable.Name,

		postgres.AccountsTable.Columns.UserID)
	err := r.db.SelectContext(ctx, &ret, query, userID)
	if err != nil {
		return nil, repository.TranslateDBError(err)
	}
	log.Println(ret)
	return ret, nil
}
