package app

import (
	"github.com/brotigen23/goph-keeper/server/docs"
	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/handler/auth"
	accountHandler "github.com/brotigen23/goph-keeper/server/internal/handler/data/account"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres"
	accountService "github.com/brotigen23/goph-keeper/server/internal/service/account"
	authService "github.com/brotigen23/goph-keeper/server/internal/service/auth"
	"github.com/brotigen23/goph-keeper/server/pkg/database"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/middleware"
	"github.com/brotigen23/goph-keeper/server/pkg/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	err = db.Migrate("file://db/postgres/migration")
	if err != nil {
		logger.Error(err)
		return err
	}

	// Repos
	repoFactory := postgres.NewFactory(db.DB)

	// Servicies
	userService := authService.New(repoFactory.NewUserRepository())
	accountService := accountService.New(repoFactory.NewAccountRepository())
	//Middleware
	middleware := middleware.New(logger, config.JWT.AccessKey, config.JWT.RefreshKey)
	// Handlers
	authHandler := auth.New(userService, config.JWT.AccessKey, config.JWT.RefreshKey)
	accountHandler := accountHandler.New(accountService)
	r := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//! ***************************************
	//! ***************************************
	//! * Routes
	//! ***************************************
	//! ***************************************

	// ***************************************
	// * Auth
	// ***************************************
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// ***************************************
	// * Users
	// ***************************************
	userGroup := r.Group("/user")
	userGroup.Use(middleware.Auth())

	// ***************************************
	// * Accounts data
	// ***************************************
	accountsGroup := userGroup.Group("/accounts")
	accountsGroup.POST("/", accountHandler.Create)
	accountsGroup.PUT("/", accountHandler.Update)
	accountsGroup.GET("/fetch", accountHandler.Fetch)

	// ***************************************
	// * Start server
	// ***************************************
	server := server.New(r, logger).Testing()
	err = server.Start()

	return err
}
