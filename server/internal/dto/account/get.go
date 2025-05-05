package account

import (
	"github.com/brotigen23/goph-keeper/server/internal/dto"
)

type GetResponse struct {
	dto.BaseData
	Model
} //@name Account.Get.Response
