package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

var textDataTable = struct {
	tableName           string
	idColumnName        string
	userIDColumnName    string
	dataColumnName      string
	createdAtColumnName string
	updatedAtColumnName string
}{"text_data", "id", "user_id", "data", "created_at", "updated_at"}

type textDataRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewTextDataRepository(db *sql.DB, logger *logger.Logger) repository.Text {
	return &textDataRepository{
		db:     db,
		logger: logger}
}

func (r textDataRepository) Create(ctx context.Context, userID int, data string) (*model.TextData, error) {

	ret := &model.TextData{
		UserID: userID,
		Data:   data,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2) RETURNING %s, %s",
		textDataTable.tableName,
		textDataTable.userIDColumnName,
		textDataTable.dataColumnName,
		textDataTable.idColumnName,
		textDataTable.createdAtColumnName)

	err = tx.QueryRowContext(ctx, query, userID, data).Scan(&ret.ID, &ret.CreatedAt)
	if err != nil {
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

func (r textDataRepository) GetByID(ctx context.Context, id int) (*model.TextData, error) {

	ret := &model.TextData{ID: id}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1",
		textDataTable.userIDColumnName,
		textDataTable.dataColumnName,
		textDataTable.createdAtColumnName,
		textDataTable.updatedAtColumnName,
		textDataTable.tableName,
		textDataTable.idColumnName)

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&ret.UserID,
			&ret.Data,
			&ret.CreatedAt,
			&ret.UpdatedAt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("text data not found", "id", id)
		return nil, repository.ErrTextDataNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}
func (r textDataRepository) GetByUserID(ctx context.Context, userID int) ([]model.TextData, error) {

	ret := []model.TextData{}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1",
		textDataTable.idColumnName,
		textDataTable.dataColumnName,
		textDataTable.createdAtColumnName,
		textDataTable.updatedAtColumnName,
		textDataTable.tableName,
		textDataTable.userIDColumnName)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	for rows.Next() {
		ret = append(ret, model.TextData{UserID: userID})
		err = rows.Scan(
			&ret[len(ret)-1].ID,
			&ret[len(ret)-1].Data,
			&ret[len(ret)-1].CreatedAt,
			&ret[len(ret)-1].UpdatedAt,
		)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
	}
	if len(ret) == 0 {
		return nil, repository.ErrTextDataNotFound
	}
	return ret, nil
}

func (r textDataRepository) Update(context.Context, model.TextData) (*model.TextData, error) {
	return nil, nil
}

func (r textDataRepository) DeleteByID(context.Context, int) (*model.TextData, error) {
	return nil, nil
}
