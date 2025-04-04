package client

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func New(baseURL string) *Client {
	client := resty.New().SetBaseURL(baseURL)
	return &Client{
		client: client,
	}
}

func (s Client) Ping() error {
	_, err := s.client.R().Get("/ping")
	return err
}
