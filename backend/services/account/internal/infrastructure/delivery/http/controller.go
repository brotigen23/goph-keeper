package http

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/usecase"
	"github.com/brotigen23/goph-keeper/accounts/internal/infrastructure/delivery/http/request"
	"github.com/brotigen23/goph-keeper/accounts/internal/infrastructure/delivery/http/response"
	"github.com/gin-gonic/gin"
)

type controller struct {
	service usecase.Usecase
}

func newHandler(service usecase.Usecase) *controller {
	return &controller{
		service: service,
	}
}

func (ctrl *controller) create(ctx *gin.Context) {
	var r request.PostAccount
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	input := usecase.Account{
		Login:    r.Login,
		Password: r.Password,
	}
	newAccount, err := ctrl.service.Create(ctx.Request.Context(), input)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	response := response.AccountPost{
		Account: response.Account{
			ID:    newAccount.ID,
			Login: newAccount.Login,
		},
	}
	ctx.JSON(http.StatusAccepted, response)
}

func (ctrl *controller) get(ctx *gin.Context) {
	filter := request.FilterParam{}
	err := ctx.ShouldBindQuery(&filter)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	usecaseFilter := usecase.ListFilter{
		Login: filter.Login,
	}
	accounts, err := ctrl.service.List(ctx.Request.Context(), usecaseFilter)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, accounts)
}
