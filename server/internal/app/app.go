package app

import (
	"context"
	"fmt"

	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/internal/mapper"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/service"
	"github.com/brotigen23/goph-keeper/server/pkg/database"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/middleware"
	"github.com/brotigen23/goph-keeper/server/pkg/server"
	"github.com/go-chi/chi/v5"
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

	serviceAggregator := service.NewAggregator(
		userRepo,
		accountsRepo,
		textDataRepo,
		binaryDataRepo,
		cardsRepo,
		metadataRepo,
	)
	acc, metadata, err := serviceAggregator.GetUserAccountsData(context.Background(), 1)
	if err != nil {
		logger.Error(err)
	}
	dto := mapper.AccountsToDTO(acc, metadata)
	fmt.Println(dto)
	//Handler
	middleware := middleware.New(logger, config.JWT.AccessKey, config.JWT.RefreshKey)
	handler := handler.New(config, serviceAggregator)

	// Router
	router := chi.NewRouter()
	router.Use(middleware.Log)

	router.Get("/ping", handler.Ping)

	// Wihtout auth
	router.Group(func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})

	// With auth
	router.Route("/user", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/accounts", handler.AccountsDataGet)
		r.Get("/text", nil)
		r.Get("/binary", nil)
		r.Get("/cards", nil)
	})

	// Server
	server := server.New(router, logger).Testing()
	err = server.Start()
	if err != nil {
		return err
	}
	return err
}
