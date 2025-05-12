package app

import (
	"fmt"
	"os"

	"github.com/brotigen23/goph-keeper/client/internal/cli/cmd"
	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
)

type App struct {
	logger *logger.Logger
}

func New() *App {
	return &App{

		logger: logger.New().Testing(),
	}
}

func (a App) Run() {
	a.logger.Info("Starting cli...")
	file, err := os.ReadFile(".cred")
	if err != nil {
		fmt.Println(err)
	}
	jwt := string(file)
	client := api.New("http://localhost:8080", jwt)
	cmd.Init(client)
	cmd.Execute()
}
