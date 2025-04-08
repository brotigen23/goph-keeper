package app

import (
	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/server"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/database"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() error {
	logger := logger.New().Testing()
	config, err := config.New().Default()
	if err != nil {
		logger.Error(err)
	}
	logger.Info("configuration", "config", config)

	//DB
	db, err := database.New("pgx", config.GetPostgresDSN())
	if err != nil {
		logger.Error(err)
		return err
	}
	defer db.Close()
	err = db.Migrate("file://migration")
	if err != nil {
		logger.Error(err)
		return err
	}

	// Repos
	// userRepo
	userRepo := postgres.NewUsersRepository(db.DB, logger)
	accountsRepo := postgres.NewAccountsRepository(db.DB, logger)
	textDataRepo := postgres.NewTextDataRepository(db.DB, logger)
	binaryDataRepo := postgres.NewBinaryRepository(db.DB, logger)
	cardsRepo := postgres.NewCardsRepository(db.DB, logger)
	metadataRepo := postgres.NewMetadataRepository(db.DB, logger)
	// Servicies
	userService := service.NewUserService(userRepo)
	accountsService := service.NewAccountsService(accountsRepo)
	textDataService := service.NewTextDataService(textDataRepo)
	binaryDataService := service.NewBinaryDataService(binaryDataRepo)
	cardsService := service.NewCardsService(cardsRepo)
	metadataService := service.NewMetadataService(metadataRepo)

	serviceAggregator := service.NewAggregator(
		userService,
		accountsService,
		textDataService,
		binaryDataService,
		cardsService,
		metadataService,
	)
	//Handler
	handler := handler.New(config, serviceAggregator)

	// Server
	middleware := middleware.New(logger, config.JWT.AccessKey, config.JWT.RefreshKey)
	server := server.New(handler, middleware, logger)

	return server.Run()
}
