package middleware

import (
	"strings"

	"github.com/brotigen23/goph-keeper/server/pkg/auth"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	logger *logger.Logger

	accessKey  string
	refreshKey string
}

func New(logger *logger.Logger, accessKey, refreshKey string) *Middleware {
	return &Middleware{
		logger:     logger,
		accessKey:  accessKey,
		refreshKey: refreshKey,
	}
}

func (m Middleware) Log() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": ErrTokenIsInvalid})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.AbortWithStatusJSON(401, gin.H{"error": ErrTokenIsInvalid})
			return
		}
		// TODO: refresh tokens if valid
		id, err := auth.GetIDFromJWT(token, m.accessKey)
		switch err {
		case nil:
			break
		case auth.ErrTokenIsInvalid:
			c.AbortWithStatusJSON(401, gin.H{"error": ErrTokenIsInvalid})
			return
		default:
			c.AbortWithStatusJSON(401, gin.H{"error": ErrTokenIsInvalid})
			return
		}
		c.Set("userID", id)
		c.Next()
	}
}
