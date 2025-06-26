package http

import (
	"github.com/brotigen23/goph-keeper/accounts/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

func AddRouterGroup(rg *gin.RouterGroup, service usecase.Usecase) error {
	handler := newHandler(service)
	r := rg.Group("/account")
	r.OPTIONS("/", nil)

	r.POST("/", handler.create)
	r.GET("/", handler.get)

	return nil
}
