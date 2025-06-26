package app

import (
	"log"

	"github.com/brotigen23/goph-keeper/accounts/internal/domain/usecase/account"
	"github.com/brotigen23/goph-keeper/accounts/internal/infrastructure/delivery/http"
	"github.com/brotigen23/goph-keeper/accounts/internal/infrastructure/repo/memory"
	"github.com/gin-gonic/gin"
)

func Run() error {
	repo := memory.NewMemory()
	service := account.New(repo)

	r := gin.Default()
	v1 := r.Group("/")
	http.AddRouterGroup(v1, service)

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
	return err
}
