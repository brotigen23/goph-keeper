package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/brotigen23/goph-keeper/backend/pkg/logger"
)

type Server struct {
	server  *http.Server
	handler http.Handler
	logger  *logger.Logger
}

func New(handler http.Handler, logger *logger.Logger) *Server {
	return &Server{
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Testing() *Server {
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s.handler,
	}
	return s
}

func (s Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	s.Shutdown()
	uptime := time.Since(start).Seconds()
	s.logger.Info("uptime", "time", uptime)
	return nil
}

func (s Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
