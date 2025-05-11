package app

import (
	"flag"

	cli "github.com/brotigen23/goph-keeper/client/internal/cli/app"
)

type App struct{}

func New() *App {
	// Services and JWT storage
	return &App{}
}

func (a App) Run() {
	isTUI := flag.Bool("t", false, "Run in tui")
	flag.Parse()
	if *isTUI {
	} else {
		CLI := cli.New()
		CLI.Run()
	}
}
