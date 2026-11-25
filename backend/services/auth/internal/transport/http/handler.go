package http

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/jwt"
	"github.com/brotigen23/goph-keeper/auth/internal/transport/http/request"
	"github.com/gin-gonic/gin"
)

type handler struct {
	controller *controller
	jwtService jwt.Service
}

func newHandler(c *controller, jwtSrc jwt.Service) *handler {
	return &handler{
		controller: c,
		jwtService: jwtSrc,
	}
}

func (h *handler) register(ctx *gin.Context) {
	r := request.Register{}
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp, err := h.controller.register(ctx.Request.Context(), r)
	if err != nil {
		ctx.Error(err)
		return
	}
	token, err := h.jwtService.Generate(resp.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Header("JWT", token)
	ctx.Status(http.StatusAccepted)
}

func (h *handler) login(ctx *gin.Context) {
	req := request.Login{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp, err := h.controller.login(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	token, err := h.jwtService.Generate(resp.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Header("JWT", token)
	ctx.Status(http.StatusAccepted)
}
