package server

import "github.com/go-resty/resty/v2"

type Server struct {
	client *resty.Client
}

func New() *Server {
	client := resty.New()
	return &Server{
		client: client,
	}
}

func (s Server) Ping() error {
	_, err := s.client.R().Get("")
	return err
}
