package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/handler"
	"github.com/brotigen23/goph-keeper/server/pkg/logger"
	"github.com/brotigen23/goph-keeper/server/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server

	handler    *handler.Handler
	middleware *middleware.Middleware

	logger *logger.Logger
}

func New(handler *handler.Handler, middleware *middleware.Middleware, logger *logger.Logger) *Server {
	return &Server{
		handler: handler,

		logger: logger,

		middleware: middleware,
	}
}

func (s Server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := chi.NewRouter()
	router.Use(s.middleware.Log)

	router.Get("/ping", s.handler.Ping)

	// Wihtout auth
	router.Group(func(r chi.Router) {
		r.Post("/register", s.handler.Register)
		r.Post("/login", s.handler.Login)
	})

	// With auth
	router.Route("/user", func(r chi.Router) {
		r.Use(s.middleware.Auth)
		r.Get("/accounts", s.handler.AccountsDataGet)
		r.Get("/text", nil)
		r.Get("/binary", nil)
		r.Get("/cards", nil)
	})

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
