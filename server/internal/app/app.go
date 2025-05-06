package app

import (
	"github.com/brotigen23/goph-keeper/server/docs"
	"github.com/brotigen23/goph-keeper/server/internal/config"
	"github.com/brotigen23/goph-keeper/server/internal/handler/auth"
	accountHandler "github.com/brotigen23/goph-keeper/server/internal/handler/data/accounthandler"
	"github.com/brotigen23/goph-keeper/server/internal/handler/data/binaryhandler"
	"github.com/brotigen23/goph-keeper/server/internal/handler/data/cardhandler"
	"github.com/brotigen23/goph-keeper/server/internal/handler/data/texthandler"
	"github.com/brotigen23/goph-keeper/server/internal/repository/postgres"
	"github.com/brotigen23/goph-keeper/server/internal/service/accountservice"
	authService "github.com/brotigen23/goph-keeper/server/internal/service/auth"
	"github.com/brotigen23/goph-keeper/server/internal/service/binaryservice"
	"github.com/brotigen23/goph-keeper/server/internal/service/cardservice"
	"github.com/brotigen23/goph-keeper/server/internal/service/textservice"
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
	accountService := accountservice.New(repoFactory.NewAccountRepository())
	textService := textservice.New(repoFactory.NewTextRepository())
	binaryService := binaryservice.New(repoFactory.NewBinaryRepository())
	cardsService := cardservice.New(repoFactory.NewCardsRepository())
	//Middleware
	middleware := middleware.New(logger, config.JWT.AccessKey, config.JWT.RefreshKey)
	// Handlers
	authHandler := auth.New(userService, config.JWT.AccessKey, config.JWT.RefreshKey)
	accountHandler := accountHandler.New(accountService)
	textHandler := texthandler.New(textService)
	binaryHandler := binaryhandler.New(binaryService)
	cardsHandler := cardhandler.New(cardsService)
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
	accountsGroup.POST("/", accountHandler.Post)
	accountsGroup.PUT("/", accountHandler.Put)
	accountsGroup.GET("/fetch", accountHandler.Fetch)
	accountsGroup.DELETE("/", accountHandler.Delete)

	// ***************************************
	// * Text data
	// ***************************************
	textGroup := userGroup.Group("/text")
	textGroup.POST("/", textHandler.Post)
	textGroup.PUT("/", textHandler.Put)
	textGroup.GET("/fetch", textHandler.Fetch)
	textGroup.DELETE("/", textHandler.Delete)

	// ***************************************
	// * Binary data
	// ***************************************
	binaryGroup := userGroup.Group("/binary")
	binaryGroup.POST("/", binaryHandler.Post)
	binaryGroup.PUT("/", binaryHandler.Put)
	binaryGroup.GET("/fetch", binaryHandler.Fetch)
	binaryGroup.DELETE("/", binaryHandler.Delete)

	// ***************************************
	// * Cards data
	// ***************************************
	cardsGroup := userGroup.Group("/cards")
	cardsGroup.POST("/", cardsHandler.Post)
	cardsGroup.PUT("/", cardsHandler.Put)
	cardsGroup.GET("/fetch", cardsHandler.Fetch)
	cardsGroup.DELETE("/", cardsHandler.Delete)

	// ***************************************
	// * Start server
	// ***************************************
	server := server.New(r, logger).Testing()
	err = server.Start()

	return err
}
