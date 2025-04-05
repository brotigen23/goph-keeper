package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/crypt"
	"github.com/brotigen23/goph-keeper/server/pkg/pgErrors"
)

var userTable = struct {
	name               string
	idColumnName       string
	loginColumnName    string
	passwordColumnName string
}{"users", "id", "login", "password"}

type UsersRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewUsers(db *sql.DB, logger *slog.Logger) *UsersRepository {
	return &UsersRepository{
		db:     db,
		logger: logger,
	}
}

func (r UsersRepository) Create(ctx context.Context, login, password string) (*model.User, error) {

	var createdAt time.Time
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2) RETURNING id, created_at",
		userTable.name,
		userTable.loginColumnName,
		userTable.passwordColumnName)

	passHash, err := crypt.HashPassword(password)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRowContext(ctx, query, login, passHash).Scan(&id, &createdAt)
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

	savedUser := &model.User{
		ID:        id,
		Login:     login,
		Password:  passHash,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	return savedUser, nil
}

func (r UsersRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}
func (r UsersRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	return nil, nil
}

func (r UsersRepository) Update(ctx context.Context, user model.User) (*model.User, error) {
	return nil, nil
}

func (r UsersRepository) DeleteByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}
