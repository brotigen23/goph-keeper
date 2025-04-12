package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
)

const cardsTableName = "cards_data"

var cardsTable = struct {
	idColumnName     string
	userIDColumnName string
	metadataID       string

	numberColumnName         string
	cardholderNameColumnName string
	expireColumnName         string
	cvvColumnName            string

	createdAtColumnName string
	updatedAtColumnName string
}{"id", "user_id", "metadata_id", "number", "cardholder_name", "expire", "cvv", "created_at", "updated_at"}

type cardsRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewCardsRepository(db *sql.DB, logger *logger.Logger) repository.Data[model.CardData] {
	return &cardsRepository{
		db:     db,
		logger: logger}
}

func (r cardsRepository) Create(ctx context.Context, data model.CardData) (*model.CardData, error) {

	ret := &model.CardData{
		UserID:         data.UserID,
		Number:         data.Number,
		CardholderName: data.CardholderName,
		Expire:         data.Expire,
		CVV:            data.CVV,
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5) RETURNING %s, %s, %s",
		cardsTableName,
		cardsTable.userIDColumnName,
		cardsTable.numberColumnName,
		cardsTable.cardholderNameColumnName,
		cardsTable.expireColumnName,
		cardsTable.cvvColumnName,
		cardsTable.idColumnName,
		cardsTable.metadataID,
		cardsTable.createdAtColumnName)

	err = tx.QueryRowContext(ctx, query,
		data.UserID, data.Number, data.CardholderName, data.Expire, data.CVV).
		Scan(&ret.ID, &ret.MetadataID, &ret.CreatedAt)
	if err != nil {
		fmt.Println(err)
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

func (r cardsRepository) GetByID(ctx context.Context, id int) (*model.CardData, error) {

	ret := &model.CardData{ID: id}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		cardsTable.userIDColumnName,
		cardsTable.numberColumnName,
		cardsTable.cardholderNameColumnName,
		cardsTable.expireColumnName,
		cardsTable.cvvColumnName,
		cardsTable.createdAtColumnName,
		cardsTable.updatedAtColumnName,
		cardsTableName,
		cardsTable.idColumnName)

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&ret.UserID,
			&ret.Number,
			&ret.CardholderName,
			&ret.Expire,
			&ret.CVV,
			&ret.CreatedAt,
			&ret.UpdatedAt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		r.logger.Info("card info not found", "id", id)
		// TODO: return error
		return nil, repository.ErrCardsDataNotFound
	default:
		r.logger.Error(err)
		return nil, err
	}

	return ret, nil
}
func (r cardsRepository) GetByUserID(ctx context.Context, userID int) ([]model.CardData, error) {

	ret := []model.CardData{}

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1",
		cardsTable.idColumnName,
		cardsTable.numberColumnName,
		cardsTable.cardholderNameColumnName,
		cardsTable.expireColumnName,
		cardsTable.cvvColumnName,
		cardsTable.createdAtColumnName,
		cardsTable.updatedAtColumnName,
		cardsTableName,
		cardsTable.userIDColumnName)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	for rows.Next() {
		ret = append(ret, model.CardData{UserID: userID})
		err = rows.Scan(
			&ret[len(ret)-1].ID,
			&ret[len(ret)-1].Number,
			&ret[len(ret)-1].CardholderName,
			&ret[len(ret)-1].Expire,
			&ret[len(ret)-1].CVV,
			&ret[len(ret)-1].CreatedAt,
			&ret[len(ret)-1].UpdatedAt,
		)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
	}

	if len(ret) == 0 {
		return nil, repository.ErrCardsDataNotFound
	}
	return ret, nil
}

func (r cardsRepository) Update(context.Context, model.CardData) (*model.CardData, error) {
	return nil, nil
}

func (r cardsRepository) DeleteByID(context.Context, int) (*model.CardData, error) { return nil, nil }
