package v1

import (
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service usecase.Usecase
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) create(ctx *gin.Context) {

}
