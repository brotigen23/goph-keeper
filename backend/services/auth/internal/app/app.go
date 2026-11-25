package app

import (
	"log/slog"

	"github.com/brotigen23/goph-keeper/auth/internal/transport/http"
	"github.com/brotigen23/goph-keeper/shared/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.New()
	r.Use(
		middleware.RequestLogger(slog.Default()),
		middleware.ErrorHandler(slog.Default(), http.MapError),
		gin.Recovery(),
	)
	api := r.Group("/api")
	http.NewRoute(api)

	err := r.Run(":8080")
	return err
}
