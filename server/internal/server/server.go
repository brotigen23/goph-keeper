package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	handler *handler.Handler
	logger  *logger.Logger

	// Service
	server *http.Server
	// middleware
}

func New(logger *logger.Logger, handler *handler.Handler) *Server {
	return &Server{
		handler: handler,

		logger: logger,
	}
}

func (s Server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/ping", s.handler.Ping)

	s.server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	////////////////////////////////////////////////////
	// START
	////////////////////////////////////////////////////
	s.logger.Info("server is running")

	start := time.Now()
	go func() {
		if e := s.server.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			s.logger.Error(e)
		}
	}()

	<-ctx.Done()
	s.logger.Info("server is shutting down")

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(err)
		return err
	}
	uptime := time.Since(start).Seconds()
	s.logger.Info("uptime", "time", uptime)
	return nil
}
