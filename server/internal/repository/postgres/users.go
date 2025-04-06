package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/pgErrors"
)

var userTable = struct {
	tableName           string
	idColumnName        string
	loginColumnName     string
	passwordColumnName  string
	createdAtColumnName string
	updatedAtColumnName string
}{"users", "id", "login", "password", "created_at", "updated_at"}

type usersRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewUsers(db *sql.DB, logger *logger.Logger) repository.Users {
	return &usersRepository{
		db:     db,
		logger: logger,
	}
}

func (r usersRepository) Create(ctx context.Context, login, password string) (*model.User, error) {

	ret := &model.User{
		Login:    login,
		Password: password,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2) RETURNING %s, %s",
		userTable.tableName,
		userTable.loginColumnName,
		userTable.passwordColumnName,
		userTable.idColumnName,
		userTable.createdAtColumnName)

	err = tx.QueryRowContext(ctx, query, login, password).Scan(&ret.ID, &ret.CreatedAt)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, rollbackErr
		}
		if pgErrors.CheckIfUniqueViolation(err) {
			return nil, repository.ErrUserExists
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	ret.UpdatedAt = ret.CreatedAt
	return ret, nil
}

func (r usersRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	ret := &model.User{}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		userTable.idColumnName,
		userTable.loginColumnName,
		userTable.passwordColumnName,
		userTable.createdAtColumnName,
		userTable.updatedAtColumnName,
		userTable.tableName,
		userTable.idColumnName)

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&ret.ID, &ret.Login, &ret.Password, &ret.CreatedAt, &ret.UpdatedAt)

	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("user not found", "userID", id)
		return nil, repository.ErrUserNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}

func (r usersRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	return nil, nil
}

func (r usersRepository) Update(ctx context.Context, user model.User) (*model.User, error) {
	return nil, nil
}

func (r usersRepository) DeleteByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}
