package mapper

import (
	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/model"
)

func AccountsToDTO(model []model.AccountData, metadata []model.Metadata) []dto.Account {
	ret := []dto.Account{}
	for i := range model {
		item := dto.Account{
			ID:         model[i].ID,
			UserID:     model[i].UserID,
			Login:      model[i].Login,
			Password:   model[i].Password,
			MetadataID: metadata[i].ID,
			Metadata:   metadata[i].Data,
		}
		ret = append(ret, item)
	}
	return ret
}

func AccountToDTO(model *model.AccountData) *dto.Account {
	return nil
}

func AccountToModel(dto *dto.Account) *model.AccountData {
	return &model.AccountData{
		ID:     dto.ID,
		UserID: dto.UserID,
	}
}
