package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

const metadataTableName = "metadata"

var metadataTable = struct {
	id string

	data string

	createdAt string
	updatedAt string
}{"id", "data", "created_at", "updated_at"}

type metadataRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewMetadataRepository(db *sql.DB, logger *logger.Logger) repository.Metadata {
	return &metadataRepository{
		db:     db,
		logger: logger}
}

func (r metadataRepository) Create(ctx context.Context, data string) (*model.Metadata, error) {

	ret := &model.Metadata{
		Data: data,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES($1) RETURNING %s, %s",
		metadataTableName,
		metadataTable.data,
		metadataTable.id,
		metadataTable.createdAt)

	err = tx.QueryRowContext(ctx, query, data).Scan(&ret.ID, &ret.CreatedAt)
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

	query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = $1",
		metadataTable.data,
		metadataTable.createdAt,
		metadataTable.updatedAt,
		metadataTableName,
		metadataTable.id)

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&ret.Data,
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

func (r metadataRepository) Update(context.Context, model.Metadata) (*model.Metadata, error) {
	return nil, nil
}

func (r metadataRepository) DeleteByID(context.Context, int) (*model.Metadata, error) {
	return nil, nil
}
