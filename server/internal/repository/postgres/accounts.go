package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

const accountsTableName = "accounts_data"

var accountsTable = struct {
	id         string
	userID     string
	metadataID string
	login      string
	password   string
	createdAt  string
	updatedAt  string
}{"id", "user_id", "metadata_id", "login", "password", "created_at", "updated_at"}

type accountsRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewAccountsRepository(db *sql.DB, logger *logger.Logger) repository.Data[model.AccountData] {
	return &accountsRepository{
		db:     db,
		logger: logger}
}

func (r accountsRepository) Create(ctx context.Context, data model.AccountData) (*model.AccountData, error) {
	ret := &model.AccountData{
		UserID:   data.UserID,
		Login:    data.Login,
		Password: data.Password,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(
		`INSERT INTO %s(
			%s, %s, %s
		) 
		VALUES($1, $2, $3) 
		RETURNING %s, %s, %s`,
		accountsTableName,
		accountsTable.userID, accountsTable.login,
		accountsTable.password, accountsTable.id,
		accountsTable.metadataID, accountsTable.createdAt)

	err = tx.QueryRowContext(ctx, query, data.UserID, data.Login, data.Password).
		Scan(&ret.ID, &ret.MetadataID, &ret.CreatedAt)

	if err != nil {
		r.logger.Error(err)
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, rollbackErr
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

func (r accountsRepository) GetByID(ctx context.Context, id int) (*model.AccountData, error) {
	ret := &model.AccountData{ID: id}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		accountsTable.userID,
		accountsTable.login,
		accountsTable.password,
		accountsTable.createdAt,
		accountsTable.updatedAt,
		accountsTableName,
		accountsTable.id)

	err := r.db.QueryRow(query, id).
		Scan(&ret.UserID,
			&ret.Login,
			&ret.Password,
			&ret.CreatedAt,
			&ret.UpdatedAt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("account not found", "id", id)
		return nil, repository.ErrAccountsDataNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}
func (r accountsRepository) GetByUserID(ctx context.Context, userID int) ([]model.AccountData, error) {

	ret := []model.AccountData{}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		accountsTable.id,
		accountsTable.login,
		accountsTable.password,
		accountsTable.createdAt,
		accountsTable.updatedAt,
		accountsTableName,
		accountsTable.userID)

	rows, err := r.db.Query(query, userID)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	for rows.Next() {
		ret = append(ret, model.AccountData{UserID: userID})
		err = rows.Scan(
			&ret[len(ret)-1].ID,
			&ret[len(ret)-1].Login,
			&ret[len(ret)-1].Password,
			&ret[len(ret)-1].CreatedAt,
			&ret[len(ret)-1].UpdatedAt,
		)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
	}

	if len(ret) == 0 {
		return nil, repository.ErrAccountsDataNotFound
	}
	return ret, nil
}

func (r accountsRepository) Update(context.Context, model.AccountData) (*model.AccountData, error) {
	return nil, nil
}

func (r accountsRepository) DeleteByID(context.Context, int) (*model.AccountData, error) {
	return nil, nil
}
