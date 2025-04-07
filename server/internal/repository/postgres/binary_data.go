package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

var binaryTable = struct {
	tableName string

	idColumnName     string
	userIDColumnName string

	dataColumnName string

	createdAtColumnName string
	updatedAtColumnName string
}{"binary_data", "id", "user_id", "data", "created_at", "updated_at"}

type binaryRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewBinaryRepository(db *sql.DB, logger *logger.Logger) repository.Binary {
	return &binaryRepository{
		db:     db,
		logger: logger}
}

func (r binaryRepository) Create(ctx context.Context, userID int, data []byte) (*model.BinaryData, error) {

	ret := &model.BinaryData{
		UserID: userID,
		Data:   data,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2) RETURNING %s, %s",
		binaryTable.tableName,
		binaryTable.userIDColumnName,
		binaryTable.dataColumnName,
		binaryTable.idColumnName,
		binaryTable.createdAtColumnName)

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

func (r binaryRepository) GetByID(ctx context.Context, id int) (*model.BinaryData, error) {

	ret := &model.BinaryData{ID: id}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1",
		binaryTable.userIDColumnName,
		binaryTable.dataColumnName,
		binaryTable.createdAtColumnName,
		binaryTable.updatedAtColumnName,
		binaryTable.tableName,
		binaryTable.idColumnName)

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
		// TODO: return error
		return nil, repository.ErrUserNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}
func (r binaryRepository) GetByUserID(ctx context.Context, userID int) ([]model.BinaryData, error) {

	ret := []model.BinaryData{}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1",
		binaryTable.idColumnName,
		binaryTable.dataColumnName,
		binaryTable.createdAtColumnName,
		binaryTable.updatedAtColumnName,
		binaryTable.tableName,
		binaryTable.userIDColumnName)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	for rows.Next() {
		ret = append(ret, model.BinaryData{UserID: userID})
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

	return ret, nil
}

func (r binaryRepository) Update(context.Context, model.BinaryData) (*model.BinaryData, error) {
	return nil, nil
}

func (r binaryRepository) DeleteByID(context.Context, int) (*model.BinaryData, error) {
	return nil, nil
}
