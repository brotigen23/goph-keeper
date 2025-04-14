package mapper

import (
	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/model"
)

func AccountsToDTO(model []model.AccountData, metadata []model.Metadata) []dto.AccountsGet {
	ret := []dto.AccountsGet{}
	for i := range model {
		met := MetadataToDTO(metadata[i])
		item := dto.AccountsGet{
			BaseData: dto.BaseData{
				ID:        model[i].ID,
				CreatedAt: model[i].CreatedAt,
				UpdatedAt: model[i].UpdatedAt,
			},
			AccountData: dto.AccountData{
				Login:    model[i].Login,
				Password: model[i].Password,
			},
			Metadata: *met,
		}
		ret = append(ret, item)
	}
	return ret
}

func AccountPutToDTO(model *model.AccountData) *dto.AccountPut {
	return &dto.AccountPut{
		BaseData: dto.BaseData{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		Login:    &model.Login,
		Password: &model.Password,
	}
}

func AccountToDTO(model *model.AccountData) *dto.AccountPost {
	return &dto.AccountPost{
		BaseData: dto.BaseData{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		AccountData: dto.AccountData{
			Login:    model.Login,
			Password: model.Password,
		},
	}
}
