package app

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/server"
	"github.com/brotigen23/goph-keeper/server/pkg/migration"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() error {
	// Logger
	logFile, err := createOrOpenLogFile(logPath)
	if err != nil {
		return err
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	// Config
	err = config.LoadDotEnv()
	if err != nil {
		logger.Error("config error", "error", err)
	}
	config := &config.Config{}
	err = config.Load()
	if err != nil {
		logger.Error("config error", "error", err)
	}
	logger.Info("configuration", "config", config)

	//DB
	db, err := sql.Open("pgx", config.GetPostgresDSN())
	if err != nil {
		logger.Error("db error", "error", err)
		return err
	}

	err = migration.Migrate(db, "file://migration/")

	// Repos
	// userRepo

	if err != nil {
		logger.Error("db error", "error", err)
		return err
	}

	// Server
	server := server.New(logger)
	return server.Run()
}

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
