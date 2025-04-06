package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

var metadataTable = struct {
	tableName string

	idColumnName        string
	tableNameColumnName string
	rowIDColumnName     string

	dataColumnName string

	createdAtColumnName string
	updatedAtColumnName string
}{"metadata", "id", "table_name", "row_id", "data", "created_at", "updated_at"}

type metadataRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewMetadataRepository(db *sql.DB, logger *logger.Logger) repository.Metadata {
	return &metadataRepository{
		db:     db,
		logger: logger}
}

func (r metadataRepository) Create(ctx context.Context, tableName string, rowID int, data string) (*model.Metadata, error) {

	ret := &model.Metadata{
		TableName: tableName,
		RowID:     rowID,
		Data:      data,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES($1, $2, $3) RETURNING %s, %s",
		metadataTable.tableName,
		metadataTable.tableNameColumnName,
		metadataTable.rowIDColumnName,
		metadataTable.dataColumnName,
		metadataTable.idColumnName,
		metadataTable.createdAtColumnName)

	err = tx.QueryRowContext(ctx, query, tableName, rowID, data).Scan(&ret.ID, &ret.CreatedAt)
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

func (r metadataRepository) GetByID(ctx context.Context, id int) (*model.Metadata, error) {

	ret := &model.Metadata{ID: id}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		metadataTable.tableNameColumnName,
		metadataTable.rowIDColumnName,
		metadataTable.dataColumnName,
		metadataTable.createdAtColumnName,
		metadataTable.updatedAtColumnName,
		metadataTable.tableName,
		metadataTable.idColumnName)

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&ret.TableName,
			&ret.RowID,
			&ret.Data,
			&ret.CreatedAt,
			&ret.UpdatedAt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("text data not found", "id", id)
		// TODO: return error
		return nil, repository.ErrMetadataNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}
func (r metadataRepository) GetByRowID(ctx context.Context, tableName string, rowID int) (*model.Metadata, error) {

	ret := &model.Metadata{TableName: tableName, RowID: rowID}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1 AND %s = $2",
		metadataTable.idColumnName,
		metadataTable.dataColumnName,
		metadataTable.createdAtColumnName,
		metadataTable.updatedAtColumnName,
		metadataTable.tableName,
		metadataTable.tableNameColumnName,
		metadataTable.rowIDColumnName)

	err := r.db.QueryRowContext(ctx, query, tableName, rowID).
		Scan(&ret.ID,
			&ret.Data,
			&ret.CreatedAt,
			&ret.UpdatedAt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("metadata not found", "row id", rowID)
		return nil, repository.ErrMetadataNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}

func (r metadataRepository) Update(context.Context, model.Metadata) (*model.Metadata, error) {
	return nil, nil
}

func (r metadataRepository) DeleteByID(context.Context, int) (*model.Metadata, error) {
	return nil, nil
}
