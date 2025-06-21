package http

import (
	"github.com/brotigen23/goph-keeper/backend/services/account/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service usecase.Usecase
}

func newHandler(service usecase.Usecase) *handler {
	return &handler{
		service: service,
	}
}

func (c *handler) create(ctx *gin.Context) {

}

func NewAccountRouter(r *gin.RouterGroup, service usecase.Usecase) error {
	// Middleware
	handler := newHandler(service)
	// Handlers
	ar := r.Group("account")
	ar.GET("list", handler.create)
	return nil
}
