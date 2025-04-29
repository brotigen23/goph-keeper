package rest

import (
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Body       []byte
	StatusCode int
	Err        error
}

type Client struct {
	client *resty.Client
	jwt    string
}

func New(baseURL string) *Client {
	client := resty.New().
		SetBaseURL(baseURL)

	return &Client{
		client: client,
	}
}

func (c Client) Ping() Response {
	response, err := c.client.R().Get("/ping")
	return Response{
		Body:       response.Body(),
		StatusCode: response.StatusCode(),
		Err:        err,
	}
}
