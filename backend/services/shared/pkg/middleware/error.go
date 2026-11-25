package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(
	logger *slog.Logger,
	fMapper func(error) (int, any),
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) == 0 {
			return
		}
		err := ctx.Errors.Last().Err
		logger.Error(
			"error",
			"method", ctx.Request.Method,
			"path", ctx.FullPath(),
			"desc", err.Error(),
		)
		code, body := fMapper(err)
		ctx.JSON(code, body)
	}
}
