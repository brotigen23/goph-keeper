package http

import (
	"github.com/gin-gonic/gin"
)

type AuthModule struct {
	Controller *controller
}

func NewRoute(rg *gin.RouterGroup) {
	h := WireInitHandler()
	authGroup := rg.Group("/auth")
	authGroup.POST("/register", h.register)
	authGroup.POST("/login", h.login)
}
