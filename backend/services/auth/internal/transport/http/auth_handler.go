package http

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/request"
	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/response"
	"github.com/gin-gonic/gin"
)

type handler struct {
	controller *controller
}

func newHandler(c *controller) *handler {
	return &handler{
		controller: c,
	}
}

func (h *handler) register(ctx *gin.Context) {
	r := request.Register{}
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = h.controller.register(ctx.Request.Context(), r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err{Msg: err.Error()})
		return
	}
	// jwt
	ctx.JSON(http.StatusAccepted, nil)
}

func (h *handler) login(ctx *gin.Context) {
	req := request.Login{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = h.controller.login(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	// jwt
	ctx.JSON(http.StatusAccepted, nil)
}
