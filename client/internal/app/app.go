package app

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/brotigen23/goph-keeper/client/internal/client"
	"github.com/brotigen23/goph-keeper/client/internal/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

var (
	logPath     = filepath.Join(".", "log")
	logFileName = logPath + "/" + time.Now().String() + ".log"
)

func createOrOpenLogFile(path string) (*os.File, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func Run() {
	logFile, err := createOrOpenLogFile(logPath)
	if err != nil {
		log.Error("cannot create or open log file", "err:", err)
		return
	}
	defer logFile.Close()

	logger := log.New(os.Stderr)
	logger.SetLevel(log.DebugLevel)
	logger.SetOutput(logFile)

	logger.Info(
		"OS INFO:",
		"OS:", runtime.GOOS,
		"Arch:", runtime.GOARCH,
		"TERM:", os.Getenv("TERM"))
	client := client.New("http://localhost:8080")
	// Main
	mainModel := model.NewMainModel(logger, client)

	p := tea.NewProgram(mainModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		logger.Error("error", "err", err)
		os.Exit(1)
	}
}
