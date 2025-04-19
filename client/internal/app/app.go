package app

import (
	"os"
	"runtime"

	"github.com/brotigen23/goph-keeper/client/internal/api"
	"github.com/brotigen23/goph-keeper/client/internal/ui/app"
	"github.com/brotigen23/goph-keeper/client/pkg/logger"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	logger := logger.New().Default()

	logger.Info(
		"OS INFO:",
		"OS:", runtime.GOOS,
		"Arch:", runtime.GOARCH,
		"TERM:", os.Getenv("TERM"))
	client := api.New("http://localhost:8080")
	// Main
	rootModel := app.New(logger, client)

	p := tea.NewProgram(rootModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
