package text

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

func New(db *sqlx.DB) repository.Text {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, item *model.TextData) error {
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES(:%s, :%s, :%s) RETURNING %s, %s",
		postgres.TextTable.Name,

		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.UserID,
		postgres.TextTable.Columns.Metadata,

		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.UserID,
		postgres.TextTable.Columns.Metadata,

		postgres.TextTable.Columns.ID,
		postgres.TextTable.Columns.CreatedAt)
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
	item.UpdatedAt = item.CreatedAt
	return nil
}

func (r *Repository) Get(ctx context.Context, id int) (*model.TextData, error) {
	ret := &model.TextData{}
	ret.ID = id
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s FROM %s WHERE %s = :%s",
		postgres.TextTable.Columns.ID,
		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.UserID,
		postgres.TextTable.Columns.Metadata,
		postgres.TextTable.Columns.CreatedAt,
		postgres.TextTable.Columns.UpdatedAt,

		postgres.TextTable.Name,

		postgres.TextTable.Columns.ID,
		postgres.TextTable.Columns.ID,
	)
	rows, err := r.db.NamedQueryContext(ctx, query, ret)
	if err != nil {
		return nil, repository.TranslateDBError(err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(ret); err != nil {
			return nil, repository.TranslateDBError(err)
		}
	}
	return ret, nil
}

func (r *Repository) Update(ctx context.Context, item *model.TextData) error {
	// Igrones created_at, updated_at and user_id
	query := fmt.Sprintf("UPDATE %s SET %s = :%s, %s = :%s WHERE %s = :%s",
		postgres.TextTable.Name,
		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.Metadata,
		postgres.TextTable.Columns.Metadata,
		postgres.TextTable.Columns.ID,
		postgres.TextTable.Columns.ID,
	)
	result, err := r.db.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if ra < 1 {
		return repository.ErrNoUpdate
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1",
		postgres.TextTable.Name,
		postgres.TextTable.Columns.ID,
	)
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context, userID int) ([]model.TextData, error) {
	ret := []model.TextData{}
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		postgres.TextTable.Columns.ID,
		postgres.TextTable.Columns.Data,
		postgres.TextTable.Columns.UserID,
		postgres.TextTable.Columns.Metadata,
		postgres.TextTable.Columns.CreatedAt,
		postgres.TextTable.Columns.UpdatedAt,

		postgres.TextTable.Name,

		postgres.TextTable.Columns.UserID)
	err := r.db.SelectContext(ctx, &ret, query, userID)
	if err != nil {
		return nil, repository.TranslateDBError(err)
	}
	log.Println(ret)
	return ret, nil
}
