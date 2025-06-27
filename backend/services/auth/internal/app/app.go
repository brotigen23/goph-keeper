package app

import (
	"github.com/brotigen23/goph-keeper/auth/internal/transport/http"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.New()
	api := r.Group("/api")
	http.NewRoute(api)

	err := r.Run(":8080")
	return err
}
