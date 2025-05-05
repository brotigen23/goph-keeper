package user

import (
	"context"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/db/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(ctx context.Context, user *model.User) error {
	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES(:%s, :%s) RETURNING %s, %s",
		postgres.UsersTable.Name,
		postgres.UsersTable.Columns.Login,
		postgres.UsersTable.Columns.Password,
		postgres.UsersTable.Columns.Login,
		postgres.UsersTable.Columns.Password,
		postgres.UsersTable.Columns.ID,
		postgres.UsersTable.Columns.CreatedAt)
	rows, err := r.db.NamedQueryContext(ctx, query, user)
	if err != nil {
		return repository.TranslateDBError(err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(user); err != nil {
			return repository.TranslateDBError(err)
		}
	}
	return nil
}

func (r UserRepository) Get(ctx context.Context, id int) (*model.User, error) {
	return nil, repository.ErrNotImplement
}

func (r UserRepository) Update(ctx context.Context, user *model.User) error {
	return repository.ErrNotImplement
}

func (r UserRepository) Delete(ctx context.Context, id int) error {
	return repository.ErrNotImplement
}

func (r UserRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	ret := &model.User{}
	query := fmt.Sprintf(`
	SELECT %s, %s, %s, %s, %s 
	FROM %s 
	WHERE %s = $1 
	LIMIT 1`,
		postgres.UsersTable.Columns.ID,
		postgres.UsersTable.Columns.Login,
		postgres.UsersTable.Columns.Password,
		postgres.UsersTable.Columns.CreatedAt,
		postgres.UsersTable.Columns.UpdatedAt,
		postgres.UsersTable.Name,
		postgres.UsersTable.Columns.Login,
	)
	err := r.db.GetContext(ctx, ret, query, login)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
