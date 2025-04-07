package app

import (
	"database/sql"

	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/server"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/migration"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() error {
	logger := logger.New().Default()

	// Config
	err := config.LoadDotEnv()
	if err != nil {
		logger.Error(err)
	}
	config := &config.Config{}
	err = config.Load()
	if err != nil {
		logger.Error(err)
	}
	logger.Info("configuration", "config", config)

	//DB
	db, err := sql.Open("pgx", config.GetPostgresDSN())
	if err != nil {
		logger.Error(err)
		return err
	}
	defer db.Close()
	err = migration.Migrate(db, "file://migration/")
	if err != nil {
		logger.Error(err)
		return err
	}

	// Repos
	// userRepo
	userRepo := postgres.NewUsersRepository(db, logger)

	// Servicies
	userService := service.NewUserService(userRepo)

	//Handler
	handler := handler.New()
	handler.SetUserService(userService)

	// Server
	server := server.New(logger, handler)

	return server.Run()
}
