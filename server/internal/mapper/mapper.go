package mapper

import (
	"github.com/brotigen23/goph-keeper/server/internal/dto"
	"github.com/brotigen23/goph-keeper/server/internal/model"
)

func MetadataToDTO(metadata model.Metadata) *dto.Metadata {
	return &dto.Metadata{
		BaseData: dto.BaseData{
			ID:        metadata.ID,
			CreatedAt: metadata.CreatedAt,
			UpdatedAt: metadata.UpdatedAt,
		},
		Data: metadata.Data,
	}
}
